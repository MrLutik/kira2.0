package monitoring

import (
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/miekg/dns"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type (
	// RamUsageInfo represents information about RAM usage.
	RamUsageInfo struct {
		TotalGB float64
		FreeGB  float64
		UsedGB  float64
	}

	// DiskUsageInfo represents information about disk usage.
	DiskUsageInfo struct {
		TotalGB uint64
		FreeGB  uint64
		UsedGB  uint64
	}
)

func (m *MonitoringService) GetCPULoadPercentage() (float64, error) {
	return getCPULoadPercentage()
}

func (m *MonitoringService) GetRAMUsage() (*RamUsageInfo, error) {
	return getRAMUsage()
}

func (m *MonitoringService) GetDiskUsage() (*DiskUsageInfo, error) {
	return getDiskUsage()
}

func (m *MonitoringService) GetPublicIP() (string, error) {
	return GetPublicIP()
}

func (m *MonitoringService) GetInterfacesIP() (map[string]string, error) {
	return getInterfacesIP()
}

// getCPULoadPercentage returns the CPU usage percentage.
func getCPULoadPercentage() (float64, error) {
	log.Infoln("Getting CPU usage percentage")

	percent, err := cpu.Percent(time.Millisecond*200, false)
	if err != nil {
		log.Errorf("Can't reach the CPU info: %s", err)
		return 0, err
	}

	return percent[0], nil
}

// getRAMUsage returns the RAM usage information.
func getRAMUsage() (*RamUsageInfo, error) {
	log.Infoln("Getting RAM usage info")

	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		log.Errorf("Can't reach the RAM usage info: %s", err)
		return nil, err
	}

	return &RamUsageInfo{
		TotalGB: float64(virtualMemory.Total) / gigabyte,
		FreeGB:  float64(virtualMemory.Free) / gigabyte,
		UsedGB:  float64(virtualMemory.Used) / gigabyte,
	}, nil
}

// getDiskUsage returns the disk usage information.
func getDiskUsage() (*DiskUsageInfo, error) {
	log.Infoln("Getting disk usage info")

	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		log.Errorf("Can't reach disk usage info: %s", err)
		return nil, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	available := stat.Bavail * uint64(stat.Bsize)

	return &DiskUsageInfo{
		TotalGB: total / gigabyte,
		FreeGB:  available / gigabyte,
		UsedGB:  (total - available) / gigabyte,
	}, nil
}

// GetPublicIP retrieves the public IP address.
//
// It performs DNS queries to determine the public IP address by querying specific DNS servers.
// The method first tries to query the 'myip.opendns.com' TXT record using the 'resolver1.opendns.com' DNS server.
// If that fails, it falls back to querying the 'o-o.myaddr.l.google.com' TXT record using the 'ns1.google.com' DNS server.
// The response from the DNS query is parsed to extract the public IP address.
//
// The method returns the public IP address as a string if it is successfully retrieved.
// If there is an error during the process, an empty string and the corresponding error are returned.
//
// TODO: Additional services such as 'http://ifconfig.me', 'https://api.ipify.org/' or 'https://ipv4.icanhazip.com/'
// can be considered as future alternatives for retrieving the public IP address.
func GetPublicIP() (string, error) {
	log.Infoln("Getting public IP address")

	client := dns.Client{Net: "udp4"}

	getPublicIPFromResponse := func(r *dns.Msg) (string, error) {
		for _, ans := range r.Answer {
			switch ans := ans.(type) {
			case *dns.A:
				log.Debugf("got `dns.A`: '%v'", ans.A.String())
				return ans.A.String(), nil
			case *dns.TXT:
				log.Debugf("got `dns.TXT`: '%v'", ans.Txt[0])
				return ans.Txt[0], nil
			}
		}
		return "", fmt.Errorf("unable to extract public IP address")
	}

	queryDNS := func(qname string, qtype uint16, server string) (string, error) {
		log.Infof("Trying the query '%s' and server '%s'", qname, server)

		message := new(dns.Msg)
		message.SetQuestion(dns.Fqdn(qname), qtype)
		r, _, err := client.Exchange(message, server)
		if err != nil {
			log.Errorf("Getting public IP error: %s", err)
			return "", err
		}
		return getPublicIPFromResponse(r)
	}

	ip, err := queryDNS("myip.opendns.com.", dns.TypeA, "resolver1.opendns.com.:53")
	if err == nil {
		return ip, nil
	}

	ip, err = queryDNS("o-o.myaddr.l.google.com.", dns.TypeTXT, "ns1.google.com.:53")
	if err == nil {
		return ip, nil
	}

	return "", fmt.Errorf("can't get the public IP address")
}

// getInterfacesIP returns a map of interface names and their corresponding IP addresses.
func getInterfacesIP() (map[string]string, error) {
	log.Infoln("Getting IP addresses of interfaces")

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Errorf("Getting interfaces info error: %s", err)
		return nil, err
	}

	interfacesIPmap := make(map[string]string, len(interfaces))

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			log.Errorf("Getting unicast interface addresses for a specific interface error: %s", err)
			return nil, err
		}

		log.Infof("Found interface: %s", iface.Name)

		ipNet, ok := addrs[0].(*net.IPNet)
		if !ok {
			continue
		}
		interfacesIPmap[iface.Name] = ipNet.IP.String()
	}

	return interfacesIPmap, nil
}
