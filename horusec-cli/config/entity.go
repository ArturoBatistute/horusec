// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//nolint
package config

import (
	"github.com/ZupIT/horusec/horusec-cli/internal/entities/toolsconfig"
	"github.com/ZupIT/horusec/horusec-cli/internal/entities/workdir"
)

const (
	// This setting has the purpose of identifying where the url where the horusec-api service is hosted will be
	// By default is http://0.0.0.0:8000
	// Validation: It is mandatory to be a valid url
	EnvHorusecAPIUri = "HORUSEC_CLI_HORUSEC_API_URI"
	// This setting will identify how long I want to wait in seconds to send the analysis object to horusec-api
	// By default is 300
	// Validation: It is mandatory to be greater than 10
	EnvTimeoutInSecondsRequest = "HORUSEC_CLI_TIMEOUT_IN_SECONDS_REQUEST"
	// This setting will identify how long I want to wait in seconds to carry out an analysis that includes:
	// acquiring a project, sending it to analysis containers and acquiring a response
	// By default is 600
	// Validation: It is mandatory to be greater than 10
	EnvTimeoutInSecondsAnalysis = "HORUSEC_CLI_TIMEOUT_IN_SECONDS_ANALYSIS"
	// This setting will identify how many in how many seconds
	// I want to check if my analysis is close to the timeout
	// By default is 15
	// Validation: It is mandatory to be greater than 10
	EnvMonitorRetryInSeconds = "HORUSEC_CLI_MONITOR_RETRY_IN_SECONDS"
	// This setting is to identify which repository you are analyzing from.
	// This repository is created within the horusec webapp
	// By default is 00000000-0000-0000-0000-000000000000
	// Validation: If exist It is mandatory to be valid uuid
	EnvRepositoryAuthorization = "HORUSEC_CLI_REPOSITORY_AUTHORIZATION"
	// This setting is to know what type of output you want for the analysis (text, json, sonarqube)
	// By default is text
	// Validation: It is mandatory to be in text, json, sonarqube
	EnvPrintOutputType = "HORUSEC_CLI_PRINT_OUTPUT_TYPE"
	// This setting is to know in which directory you want the output of the json file
	// generated by the output types json or sonarqube to be located.
	// By default if the type is json or sonarqube o path is ./output.json
	// Validation: It is mandatory to be valid path
	EnvJSONOutputFilePath = "HORUSEC_CLI_JSON_OUTPUT_FILEPATH"
	// This setting is to find out what types of severity I don't want you to recognize as a vulnerability.
	// The types are: "LOW", "MEDIUM", "HIGH", "NOSEC", "AUDIT"
	// If you want ignore other you can add in value. Ex.: "LOW, MEDIUM, NOSEC"
	// This setting is to know what types of severity
	// I do not want you to recognize as a vulnerability
	// and will not count towards the return of exit (1) if configured
	// Validation: It is mandatory to be in "LOW", "MEDIUM", "HIGH", "NOSEC", "AUDIT"
	EnvSeveritiesToIgnore = "HORUSEC_CLI_SEVERITIES_TO_IGNORE"
	// This setting is to know which files and folders I want to ignore to send for analysis
	// By default we ignore each other:
	//   * Folders: "/.horusec/", "/.idea/", "/.vscode/", "/tmp/", "/bin/", "/node_modules/", "/vendor/"
	//   * Files: ".jpg", ".png", ".gif", ".webp", ".tiff", ".psd", ".raw", ".bmp", ".heif", ".indd",
	//		".jpeg", ".svg", ".ai", ".eps", ".pdf", ".webm", ".mpg", ".mp2", ".mpeg", ".mpe",
	//		".mp4", ".m4p", ".m4v", ".avi", ".wmv", ".mov", ".qt", ".flv", ".swf", ".avchd", ".mpv", ".ogg",
	EnvFilesOrPathsToIgnore = "HORUSEC_CLI_FILES_OR_PATHS_TO_IGNORE"
	// This setting is to know if I want return exit(1) if I find any vulnerability in the analysis
	// By default is false
	// Validation: It is mandatory to be in "false", "true"
	EnvReturnErrorIfFoundVulnerability = "HORUSEC_CLI_RETURN_ERROR_IF_FOUND_VULNERABILITY"
	// This setting is to know if I want to change the analysis directory
	// and do not want to run in the current directory.
	// If this value is not passed, Horusec will ask if you want to run the analysis in the current directory.
	// If you pass it it will start the analysis in the directory informed by you without asking anything.
	// By default is CURRENT DIRECTORY
	// Validation: It is mandatory to be valid path
	EnvProjectPath = "HORUSEC_CLI_PROJECT_PATH"
	// This setting is to know in which directory I want to perform the analysis of each language.
	// As a key you must pass the name of the language and the value the directory from within your project.
	// Example:
	// Let's assume that your project is a netcore app using angular and has the following structure:
	// - NetCoreProject/
	//   - controllers/
	//   - NetCoreProject.csproj
	//   - views/
	//     - pages/
	//     - package.json
	//     - package-lock.json
	// Then your workdir would be:
	// {
	//   "csharp": ["NetCoreProject"],
	//   "javaScript": ["NetCoreProject/views"]
	// }
	// The interface is:
	// {
	//   go         []string
	//   netCore    []string DEPRECATED on 23 nov 2020
	//   csharp     []string
	//   ruby       []string
	//   python     []string
	//   java       []string
	//   kotlin     []string
	//   javaScript []string
	//   leaks      []string
	//   hcl        []string
	//   php        []string
	//   c          []string
	//   yaml       []string
	//   generic    []string
	// }
	// Validation: It is mandatory to be valid interface of workdir to proceed
	EnvWorkDir = "HORUSEC_CLI_WORK_DIR"
	// This setting is to setup the path to run analysis keep current path in your base.
	// By default is empty
	// Validation: if exists is required valid path
	EnvFilterPath = "HORUSEC_CLI_FILTER_PATH"
	// This setting is to know if I want enable run gitleaks tools
	// and analysis in all git history searching vulnerabilities
	// By default is false
	// Validation: It is mandatory to be in "false", "true"
	EnvEnableGitHistoryAnalysis = "HORUSEC_CLI_ENABLE_GIT_HISTORY_ANALYSIS"
	// Used to authorize the sending of unsafe requests. Its use is not recommended outside testing scenarios.
	// By default is false
	// Validation: It is mandatory to be in "false", "true"
	EnvCertInsecureSkipVerify = "HORUSEC_CLI_CERT_INSECURE_SKIP_VERIFY"
	// Used to pass the path to a certificate that will be sent on the http request to the horusec server.
	// Example: /home/certs/ca.crt
	// Validation: It must be a valid path
	EnvCertPath = "HORUSEC_CLI_CERT_PATH"
	// Used to enable or disable search with vulnerability author.
	// By default is false
	// Validation: It is mandatory to be in "false", "true"
	EnvEnableCommitAuthor = "HORUSEC_CLI_ENABLE_COMMIT_AUTHOR"
	// Used to send the repository name to the server, must be used together with the company token.
	// By default is empty
	EnvRepositoryName = "HORUSEC_CLI_REPOSITORY_NAME"
	// Used to skip vulnerability of type false positive
	// By default is empty
	EnvFalsePositiveHashes = "HORUSEC_CLI_FALSE_POSITIVE_HASHES"
	// Used to skip vulnerability of type risk accept
	// By default is empty
	EnvRiskAcceptHashes = "HORUSEC_CLI_RISK_ACCEPT_HASHES"
	// DEPRECATED on 16 dec 2020
	EnvToolsToIgnore = "HORUSEC_CLI_TOOLS_TO_IGNORE"
	// Used to set configurations of tools
	// By default is setup:
	// {
	//
	// }
	EnvToolsConfig = "HORUSEC_CLI_TOOLS_CONFIG"
	// Used send others headers on request to send in horusec-api
	// By default is empty
	EnvHeaders = "HORUSEC_CLI_HEADERS"
	// Used to pass project path in host when running horusec cli inside a container
	// By default is empty
	EnvContainerBindProjectPath = "HORUSEC_CLI_CONTAINER_BIND_PROJECT_PATH"
	// Used to run horusec without docker if enabled it will only run the following tools: horusec-csharp, horusec-kotlin, horusec-kubernetes, horusec-leaks, horusec-nodejs.
	// By default is false
	// Validation: It is mandatory to be in "false", "true"
	EnvDisableDocker = "HORUSEC_CLI_DISABLE_DOCKER"
	// Used to pass the path to the horusec custom rules file. Example: -c="./horusec/horusec-custom-rules.json".
	// By default is empty
	// Validation: It is mandatory to be a valida path and contains file name
	EnvCustomRulesPath = "HORUSEC_CLI_CUSTOM_RULES_PATH"
	// Used to enable or disable information severity vulnerabilities, information vulnerabilities can contain a lot of false positives.
	// By default is false
	// Validation: It is mandatory to be in "false", "true"
	EnvEnableInformationSeverity = "HORUSEC_CLI_ENABLE_INFORMATION_SEVERITY"
)

type Config struct {
	// Globals Command Flags
	logLevel       string
	configFilePath string

	// Start Command Flags
	horusecAPIUri                   string
	repositoryAuthorization         string
	filterPath                      string
	certPath                        string
	repositoryName                  string
	printOutputType                 string
	jsonOutputFilePath              string
	projectPath                     string
	customRulesPath                 string
	containerBindProjectPath        string
	timeoutInSecondsRequest         int64
	timeoutInSecondsAnalysis        int64
	monitorRetryInSeconds           int64
	isTimeout                       bool
	returnErrorIfFoundVulnerability bool
	enableGitHistoryAnalysis        bool
	certInsecureSkipVerify          bool
	enableCommitAuthor              bool
	disableDocker                   bool
	enableInformationSeverity       bool
	severitiesToIgnore              []string
	filesOrPathsToIgnore            []string
	falsePositiveHashes             []string
	riskAcceptHashes                []string
	toolsToIgnore                   []string
	toolsConfig                     toolsconfig.MapToolConfig
	headers                         map[string]string
	workDir                         *workdir.WorkDir
}
