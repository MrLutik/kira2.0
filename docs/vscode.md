# Visual Studio Code Setup Guide

This guide provides step-by-step instructions on how to configure Visual Studio Code (VS Code) for optimal development experience in this project.

## Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/dl/)
- [Visual Studio Code](https://code.visualstudio.com/)

## Setting Up Workspace Settings

The project contains a `.vscode` folder with custom settings for VS Code. These settings help maintain consistency in coding style and tool usage across the development team.

### Key features settings

- Language Server Protocol (LSP `gopls`) Configuration
- Sets the default formatter to golang.go, which is the official Go extension  
(tools required: `gopls`, `dlv`, `gofumpt`, `golangci-lint`)
- Enables code actions `on save`
- Turns on format `on save` mode for Go files.
- Provides inline snippet suggestions within the editor.
- Specifies [`golangci-lint`](https://golangci-lint.run/) as the linting tool, which is a comprehensive linter for Go.

### Applying Workspace Settings

1. Open the project folder in VS Code.
2. VS Code automatically detects the `.vscode` folder and applies the settings defined in `settings.json`.
3. Verify that the settings are active by checking your workspace settings. Go to `File > Preferences > Settings` and ensure they match the contents of the project's `settings.json`.

## Installing Recommended Extensions

The `extensions.json` file in the `.vscode` folder lists recommended extensions that enhance the development experience.

### Installing Extensions

1. Open the Extensions view in VS Code by clicking on the Extensions icon in the Activity Bar on the side of the window or pressing `Ctrl+Shift+X`. 
2. VS Code prompts you with a notification to install recommended extensions defined in `extensions.json`.
3. Click on "Show Recommendations" to view the list.
4. Install each recommended extension by clicking on the `Install` button beside each extension in the list.

### Extensions Included:

- **golang.go**: The official Go extension for language support, debugging, and tool integration.  
(most of the settings are based on this extension)

- **code-spell-checker**: A basic spell checker that works well with camelCase.
- **gitlens**: Enhances the built-in Git capabilities in VS Code.

## Final Steps

After setting up the workspace settings and installing the recommended extensions, your VS Code environment should be configured for optimal use with this project. 

## Optional Settings

If there are any unwanted extensions that you think should not be used in this project, list them in the "unwantedRecommendations" array in `extensions.json`. Other developers will be alerted not to use these extensions in this workspace.

---

*Note: This guide is tailored for the specific settings and extensions provided. Adjustments may be necessary if there are updates or changes in the project requirements.*
