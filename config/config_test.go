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

package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/golang/mock/gomock"

	mock_config "github.com/ZupIT/horusec/config/mocks"

	"github.com/sirupsen/logrus"

	"github.com/ZupIT/horusec-devkit/pkg/enums/vulnerability"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/ZupIT/horusec-devkit/pkg/enums/languages"
	"github.com/ZupIT/horusec-devkit/pkg/enums/tools"
	"github.com/ZupIT/horusec/internal/entities/toolsconfig"
	"github.com/ZupIT/horusec/internal/entities/workdir"
)

func TestMain(m *testing.M) {
	_ = os.RemoveAll("./tmp")
	_ = os.MkdirAll("./tmp", 0750)
	code := m.Run()
	_ = os.RemoveAll("./tmp")
	os.Exit(code)
}

func TestNewHorusecConfig(t *testing.T) {
	//	wd := &workdir.WorkDir{}
	t.Run("Should return horusec config with your default values", func(t *testing.T) {
		currentPath, _ := os.Getwd()
		configs := &Config{}
		configs.NewConfigsFromEnvironments()
		assert.Equal(t, "{{VERSION_NOT_FOUND}}", configs.GetVersion())
		assert.Equal(t, configs.GetDefaultConfigFilePath(), configs.GetConfigFilePath())
		assert.Equal(t, "http://0.0.0.0:8000", configs.GetHorusecAPIUri())
		assert.Equal(t, int64(300), configs.GetTimeoutInSecondsRequest())
		assert.Equal(t, int64(600), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, int64(15), configs.GetMonitorRetryInSeconds())
		assert.Equal(t, uuid.Nil.String(), configs.GetRepositoryAuthorization())
		assert.Equal(t, "", configs.GetPrintOutputType())
		assert.Equal(t, "", configs.GetJSONOutputFilePath())
		assert.Equal(t, 1, len(configs.GetSeveritiesToIgnore()))
		assert.Equal(t, 2, len(configs.GetFilesOrPathsToIgnore()))
		assert.Equal(t, false, configs.GetReturnErrorIfFoundVulnerability())
		assert.Equal(t, currentPath, configs.GetProjectPath())
		assert.Equal(t, Config{}.workDir, configs.GetWorkDir())
		assert.Equal(t, false, configs.GetEnableGitHistoryAnalysis())
		assert.Equal(t, false, configs.GetCertInsecureSkipVerify())
		assert.Equal(t, "", configs.GetCertPath())
		assert.Equal(t, false, configs.GetEnableCommitAuthor())
		assert.Equal(t, "config", configs.GetRepositoryName())
		assert.Equal(t, 0, len(configs.GetRiskAcceptHashes()))
		assert.Equal(t, 0, len(configs.GetFalsePositiveHashes()))
		assert.Equal(t, 0, len(configs.GetHeaders()))
		assert.Equal(t, "", configs.GetContainerBindProjectPath())
		assert.Equal(t, true, configs.IsEmptyRepositoryAuthorization())
		assert.Equal(t, 0, len(configs.GetToolsConfig()))
		assert.Equal(t, false, configs.GetDisableDocker())
		assert.Equal(t, "", configs.GetCustomRulesPath())
		assert.Equal(t, false, configs.GetEnableInformationSeverity())
		assert.Equal(t, 0, len(configs.GetCustomImages()))
		assert.Equal(t, 1, len(configs.GetShowVulnerabilitiesTypes()))
		assert.Equal(t, false, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, false, configs.GetEnableShellCheck())
	})
	t.Run("Should change horusec config and return your new values", func(t *testing.T) {
		currentPath, _ := os.Getwd()
		configs := &Config{}
		configs.SetConfigFilePath(path.Join(currentPath + "other-horusec-config.json"))
		configs.SetHorusecAPIURI(uuid.New().String())
		configs.SetTimeoutInSecondsRequest(1010)
		configs.SetTimeoutInSecondsAnalysis(1010)
		configs.SetMonitorRetryInSeconds(1010)
		configs.SetRepositoryAuthorization(uuid.New().String())
		configs.SetPrintOutputType("json")
		configs.SetJSONOutputFilePath("./other-file-path.json")
		configs.SetSeveritiesToIgnore([]string{"info"})
		configs.SetFilesOrPathsToIgnore([]string{"**/*_test.go"})
		configs.SetReturnErrorIfFoundVulnerability(true)
		configs.SetProjectPath("./some-other-file-path")
		configs.SetWorkDir(map[string]interface{}{"csharp": []string{"test"}})
		configs.SetEnableGitHistoryAnalysis(true)
		configs.SetCertInsecureSkipVerify(true)
		configs.SetCertPath("./certs")
		configs.SetEnableCommitAuthor(true)
		configs.SetRepositoryName("my-project")
		configs.SetRiskAcceptHashes([]string{"123456789"})
		configs.SetFalsePositiveHashes([]string{"987654321"})
		configs.SetHeaders(map[string]string{"x-header": "value"})
		configs.SetContainerBindProjectPath("./some-other-file-path")
		configs.SetIsTimeout(true)
		configs.SetToolsConfig(toolsconfig.MapToolConfig{tools.Sobelow: {IsToIgnore: true}})
		configs.SetDisableDocker(true)
		configs.SetCustomRulesPath("test")
		configs.SetEnableInformationSeverity(true)
		configs.SetCustomImages(map[languages.Language]string{languages.Go: "test/test"})
		configs.SetShowVulnerabilitiesTypes([]string{vulnerability.Vulnerability.ToString()})
		configs.SetEnableOwaspDependencyCheck(true)
		configs.SetEnableShellCheck(true)

		assert.NotEqual(t, configs.GetDefaultConfigFilePath(), configs.GetConfigFilePath())
		assert.NotEqual(t, "http://0.0.0.0:8000", configs.GetHorusecAPIUri())
		assert.NotEqual(t, int64(300), configs.GetTimeoutInSecondsRequest())
		assert.NotEqual(t, int64(600), configs.GetTimeoutInSecondsAnalysis())
		assert.NotEqual(t, int64(15), configs.GetMonitorRetryInSeconds())
		assert.NotEqual(t, uuid.Nil.String(), configs.GetRepositoryAuthorization())
		assert.NotEqual(t, "text", configs.GetPrintOutputType())
		assert.NotEqual(t, "", configs.GetJSONOutputFilePath())
		assert.NotEqual(t, 0, len(configs.GetSeveritiesToIgnore()))
		assert.NotEqual(t, 0, len(configs.GetFilesOrPathsToIgnore()))
		assert.NotEqual(t, false, configs.GetReturnErrorIfFoundVulnerability())
		assert.NotEqual(t, currentPath, configs.GetProjectPath())
		assert.NotEqual(t, workdir.NewWorkDir().CSharp, configs.GetWorkDir().CSharp)
		assert.NotEqual(t, false, configs.GetEnableGitHistoryAnalysis())
		assert.NotEqual(t, false, configs.GetCertInsecureSkipVerify())
		assert.NotEqual(t, "", configs.GetCertPath())
		assert.NotEqual(t, false, configs.GetEnableCommitAuthor())
		assert.NotEqual(t, "", configs.GetRepositoryName())
		assert.NotEqual(t, 0, len(configs.GetRiskAcceptHashes()))
		assert.NotEqual(t, 0, len(configs.GetFalsePositiveHashes()))
		assert.NotEqual(t, 0, len(configs.GetHeaders()))
		assert.NotEqual(t, "", configs.GetContainerBindProjectPath())
		assert.NotEqual(t, false, configs.GetIsTimeout())
		assert.NotEqual(t, false, configs.GetToolsConfig()[tools.Sobelow])
		assert.Equal(t, true, configs.GetDisableDocker())
		assert.Equal(t, "test", configs.GetCustomRulesPath())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString()}, configs.GetShowVulnerabilitiesTypes())
		assert.NotEqual(t, map[languages.Language]string{}, configs.GetCustomImages())
		assert.Equal(t, true, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, true, configs.GetEnableShellCheck())
	})
	t.Run("Should return horusec config using new viper file", func(t *testing.T) {
		viper.Reset()
		currentPath, err := os.Getwd()
		configFilePath := path.Join(currentPath + "/.example-horusec-cli.json")
		assert.NoError(t, err)
		configs := &Config{}
		configs.SetConfigFilePath(configFilePath)
		configs.NewConfigsFromViper()
		assert.Equal(t, configFilePath, configs.GetConfigFilePath())
		assert.Equal(t, "http://new-viper.horusec.com", configs.GetHorusecAPIUri())
		assert.Equal(t, int64(20), configs.GetTimeoutInSecondsRequest())
		assert.Equal(t, int64(100), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, int64(10), configs.GetMonitorRetryInSeconds())
		assert.Equal(t, "8beffdca-636e-4d73-a22f-b0f7c3cff1c4", configs.GetRepositoryAuthorization())
		assert.Equal(t, "json", configs.GetPrintOutputType())
		assert.Equal(t, "./output.json", configs.GetJSONOutputFilePath())
		assert.Equal(t, []string{"INFO"}, configs.GetSeveritiesToIgnore())
		assert.Equal(t, []string{"./assets"}, configs.GetFilesOrPathsToIgnore())
		assert.Equal(t, true, configs.GetReturnErrorIfFoundVulnerability())
		assert.Equal(t, "./", configs.GetProjectPath())
		assert.Equal(t, workdir.NewWorkDir(), configs.GetWorkDir())
		assert.Equal(t, true, configs.GetEnableGitHistoryAnalysis())
		assert.Equal(t, true, configs.GetCertInsecureSkipVerify())
		assert.Equal(t, "", configs.GetCertPath())
		assert.Equal(t, true, configs.GetEnableCommitAuthor())
		assert.Equal(t, "horus", configs.GetRepositoryName())
		assert.Equal(t, []string{"hash3", "hash4"}, configs.GetRiskAcceptHashes())
		assert.Equal(t, []string{"hash1", "hash2"}, configs.GetFalsePositiveHashes())
		assert.Equal(t, map[string]string{"x-headers": "some-other-value"}, configs.GetHeaders())
		assert.Equal(t, "test", configs.GetContainerBindProjectPath())
		assert.Equal(t, true, configs.GetDisableDocker())
		assert.Equal(t, "test", configs.GetCustomRulesPath())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
		assert.Equal(t, true, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, true, configs.GetEnableShellCheck())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString(), vulnerability.FalsePositive.ToString()}, configs.GetShowVulnerabilitiesTypes())
		assert.Equal(t, toolsconfig.ToolConfig{
			IsToIgnore: true,
		}, configs.GetToolsConfig()[tools.GoSec])
		assert.Equal(t, "docker.io/company/go:latest", configs.GetCustomImages()["go"])
	})
	t.Run("Should return horusec config using viper file and override by environment", func(t *testing.T) {
		viper.Reset()
		authorization := uuid.New().String()
		currentPath, err := os.Getwd()
		configFilePath := path.Join(currentPath + "/.example-horusec-cli.json")
		assert.NoError(t, err)
		configs := &Config{}
		configs.SetConfigFilePath(configFilePath)
		configs.NewConfigsFromViper()
		assert.Equal(t, configFilePath, configs.GetConfigFilePath())
		assert.Equal(t, "http://new-viper.horusec.com", configs.GetHorusecAPIUri())
		assert.Equal(t, int64(20), configs.GetTimeoutInSecondsRequest())
		assert.Equal(t, int64(100), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, int64(10), configs.GetMonitorRetryInSeconds())
		assert.Equal(t, "8beffdca-636e-4d73-a22f-b0f7c3cff1c4", configs.GetRepositoryAuthorization())
		assert.Equal(t, "json", configs.GetPrintOutputType())
		assert.Equal(t, "./output.json", configs.GetJSONOutputFilePath())
		assert.Equal(t, []string{"INFO"}, configs.GetSeveritiesToIgnore())
		assert.Equal(t, []string{"./assets"}, configs.GetFilesOrPathsToIgnore())
		assert.Equal(t, true, configs.GetReturnErrorIfFoundVulnerability())
		assert.Equal(t, "./", configs.GetProjectPath())
		assert.Equal(t, workdir.NewWorkDir(), configs.GetWorkDir())
		assert.Equal(t, true, configs.GetEnableGitHistoryAnalysis())
		assert.Equal(t, true, configs.GetCertInsecureSkipVerify())
		assert.Equal(t, "", configs.GetCertPath())
		assert.Equal(t, true, configs.GetEnableCommitAuthor())
		assert.Equal(t, "horus", configs.GetRepositoryName())
		assert.Equal(t, []string{"hash3", "hash4"}, configs.GetRiskAcceptHashes())
		assert.Equal(t, []string{"hash1", "hash2"}, configs.GetFalsePositiveHashes())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString(), vulnerability.FalsePositive.ToString()}, configs.GetShowVulnerabilitiesTypes())
		assert.Equal(t, map[string]string{"x-headers": "some-other-value"}, configs.GetHeaders())
		assert.Equal(t, "test", configs.GetContainerBindProjectPath())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
		assert.Equal(t, true, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, true, configs.GetEnableShellCheck())
		assert.Equal(t, toolsconfig.ToolConfig{
			IsToIgnore: true,
		}, configs.GetToolsConfig()[tools.GoSec])
		assert.Equal(t, "docker.io/company/go:latest", configs.GetCustomImages()["go"])

		assert.NoError(t, os.Setenv(EnvHorusecAPIUri, "http://horusec.com"))
		assert.NoError(t, os.Setenv(EnvTimeoutInSecondsRequest, "99"))
		assert.NoError(t, os.Setenv(EnvTimeoutInSecondsAnalysis, "999"))
		assert.NoError(t, os.Setenv(EnvMonitorRetryInSeconds, "20"))
		assert.NoError(t, os.Setenv(EnvRepositoryAuthorization, authorization))
		assert.NoError(t, os.Setenv(EnvPrintOutputType, "sonarqube"))
		assert.NoError(t, os.Setenv(EnvJSONOutputFilePath, "./output-sonarqube.json"))
		assert.NoError(t, os.Setenv(EnvSeveritiesToIgnore, "INFO"))
		assert.NoError(t, os.Setenv(EnvFilesOrPathsToIgnore, "**/*_test.go, **/*_mock.go"))
		assert.NoError(t, os.Setenv(EnvReturnErrorIfFoundVulnerability, "false"))
		assert.NoError(t, os.Setenv(EnvProjectPath, "./horusec-manager"))
		assert.NoError(t, os.Setenv(EnvEnableGitHistoryAnalysis, "false"))
		assert.NoError(t, os.Setenv(EnvCertInsecureSkipVerify, "false"))
		assert.NoError(t, os.Setenv(EnvCertPath, "./"))
		assert.NoError(t, os.Setenv(EnvEnableCommitAuthor, "false"))
		assert.NoError(t, os.Setenv(EnvRepositoryName, "my-project"))
		assert.NoError(t, os.Setenv(EnvFalsePositiveHashes, "hash9, hash8"))
		assert.NoError(t, os.Setenv(EnvRiskAcceptHashes, "hash7, hash6"))
		assert.NoError(t, os.Setenv(EnvHeaders, "{\"x-auth\": \"987654321\"}"))
		assert.NoError(t, os.Setenv(EnvContainerBindProjectPath, "./my-path"))
		assert.NoError(t, os.Setenv(EnvDisableDocker, "true"))
		assert.NoError(t, os.Setenv(EnvEnableOwaspDependencyCheck, "true"))
		assert.NoError(t, os.Setenv(EnvEnableShellCheck, "true"))
		assert.NoError(t, os.Setenv(EnvCustomRulesPath, "test"))
		assert.NoError(t, os.Setenv(EnvEnableInformationSeverity, "true"))
		assert.NoError(t, os.Setenv(EnvShowVulnerabilitiesTypes, fmt.Sprintf("%s, %s", vulnerability.Vulnerability.ToString(), vulnerability.RiskAccepted.ToString())))
		assert.NoError(t, os.Setenv(EnvLogFilePath, "test"))
		configs.NewConfigsFromEnvironments()
		assert.Equal(t, configFilePath, configs.GetConfigFilePath())
		assert.Equal(t, "http://horusec.com", configs.GetHorusecAPIUri())
		assert.Equal(t, int64(99), configs.GetTimeoutInSecondsRequest())
		assert.Equal(t, int64(999), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, int64(20), configs.GetMonitorRetryInSeconds())
		assert.Equal(t, authorization, configs.GetRepositoryAuthorization())
		assert.Equal(t, "sonarqube", configs.GetPrintOutputType())
		assert.Equal(t, "./output-sonarqube.json", configs.GetJSONOutputFilePath())
		assert.Equal(t, []string{"INFO"}, configs.GetSeveritiesToIgnore())
		assert.Equal(t, []string{"**/*_test.go", "**/*_mock.go"}, configs.GetFilesOrPathsToIgnore())
		assert.Equal(t, false, configs.GetReturnErrorIfFoundVulnerability())
		assert.Equal(t, "./horusec-manager", configs.GetProjectPath())
		assert.Equal(t, workdir.NewWorkDir(), configs.GetWorkDir())
		assert.Equal(t, false, configs.GetEnableGitHistoryAnalysis())
		assert.Equal(t, false, configs.GetCertInsecureSkipVerify())
		assert.Equal(t, "./", configs.GetCertPath())
		assert.Equal(t, false, configs.GetEnableCommitAuthor())
		assert.Equal(t, "my-project", configs.GetRepositoryName())
		assert.Equal(t, []string{"hash7", "hash6"}, configs.GetRiskAcceptHashes())
		assert.Equal(t, []string{"hash9", "hash8"}, configs.GetFalsePositiveHashes())
		assert.Equal(t, map[string]string{"x-auth": "987654321"}, configs.GetHeaders())
		assert.Equal(t, "./my-path", configs.GetContainerBindProjectPath())
		assert.Equal(t, true, configs.GetDisableDocker())
		assert.Equal(t, "test", configs.GetCustomRulesPath())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
		assert.Equal(t, true, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, true, configs.GetEnableShellCheck())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString(), vulnerability.RiskAccepted.ToString()}, configs.GetShowVulnerabilitiesTypes())
	})
	t.Run("Should return horusec config using viper file and override by environment and override by flags", func(t *testing.T) {
		viper.Reset()
		authorization := uuid.New().String()
		currentPath, err := os.Getwd()
		configFilePath := path.Join(currentPath + "/.example-horusec-cli.json")
		assert.NoError(t, err)
		configs := &Config{}
		configs.factoryParseInputToSliceString(map[string]interface{}{})
		configs.SetConfigFilePath(configFilePath)
		configs.NewConfigsFromViper()
		assert.Equal(t, configFilePath, configs.GetConfigFilePath())
		assert.Equal(t, "http://new-viper.horusec.com", configs.GetHorusecAPIUri())
		assert.Equal(t, int64(20), configs.GetTimeoutInSecondsRequest())
		assert.Equal(t, int64(100), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, int64(10), configs.GetMonitorRetryInSeconds())
		assert.Equal(t, "8beffdca-636e-4d73-a22f-b0f7c3cff1c4", configs.GetRepositoryAuthorization())
		assert.Equal(t, "json", configs.GetPrintOutputType())
		assert.Equal(t, "./output.json", configs.GetJSONOutputFilePath())
		assert.Equal(t, []string{"INFO"}, configs.GetSeveritiesToIgnore())
		assert.Equal(t, []string{"./assets"}, configs.GetFilesOrPathsToIgnore())
		assert.Equal(t, true, configs.GetReturnErrorIfFoundVulnerability())
		assert.Equal(t, "./", configs.GetProjectPath())
		assert.Equal(t, workdir.NewWorkDir(), configs.GetWorkDir())
		assert.Equal(t, true, configs.GetEnableGitHistoryAnalysis())
		assert.Equal(t, true, configs.GetCertInsecureSkipVerify())
		assert.Equal(t, "", configs.GetCertPath())
		assert.Equal(t, true, configs.GetEnableCommitAuthor())
		assert.Equal(t, "horus", configs.GetRepositoryName())
		assert.Equal(t, []string{"hash3", "hash4"}, configs.GetRiskAcceptHashes())
		assert.Equal(t, []string{"hash1", "hash2"}, configs.GetFalsePositiveHashes())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString(), vulnerability.FalsePositive.ToString()}, configs.GetShowVulnerabilitiesTypes())
		assert.Equal(t, map[string]string{"x-headers": "some-other-value"}, configs.GetHeaders())
		assert.Equal(t, "test", configs.GetContainerBindProjectPath())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
		assert.Equal(t, true, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, true, configs.GetEnableShellCheck())
		assert.Equal(t, toolsconfig.ToolConfig{
			IsToIgnore: true,
		}, configs.GetToolsConfig()[tools.GoSec])
		assert.Equal(t, "docker.io/company/go:latest", configs.GetCustomImages()["go"])

		assert.NoError(t, os.Setenv(EnvHorusecAPIUri, "http://horusec.com"))
		assert.NoError(t, os.Setenv(EnvTimeoutInSecondsRequest, "99"))
		assert.NoError(t, os.Setenv(EnvTimeoutInSecondsAnalysis, "999"))
		assert.NoError(t, os.Setenv(EnvMonitorRetryInSeconds, "20"))
		assert.NoError(t, os.Setenv(EnvRepositoryAuthorization, authorization))
		assert.NoError(t, os.Setenv(EnvPrintOutputType, "sonarqube"))
		assert.NoError(t, os.Setenv(EnvJSONOutputFilePath, "./output-sonarqube.json"))
		assert.NoError(t, os.Setenv(EnvSeveritiesToIgnore, "INFO"))
		assert.NoError(t, os.Setenv(EnvFilesOrPathsToIgnore, "**/*_test.go, **/*_mock.go"))
		assert.NoError(t, os.Setenv(EnvReturnErrorIfFoundVulnerability, "false"))
		assert.NoError(t, os.Setenv(EnvProjectPath, "./horusec-manager"))
		assert.NoError(t, os.Setenv(EnvEnableGitHistoryAnalysis, "false"))
		assert.NoError(t, os.Setenv(EnvCertInsecureSkipVerify, "false"))
		assert.NoError(t, os.Setenv(EnvCertPath, "./"))
		assert.NoError(t, os.Setenv(EnvEnableCommitAuthor, "false"))
		assert.NoError(t, os.Setenv(EnvRepositoryName, "my-project"))
		assert.NoError(t, os.Setenv(EnvFalsePositiveHashes, "hash9, hash8"))
		assert.NoError(t, os.Setenv(EnvRiskAcceptHashes, "hash7, hash6"))
		assert.NoError(t, os.Setenv(EnvHeaders, "{\"x-auth\": \"987654321\"}"))
		assert.NoError(t, os.Setenv(EnvContainerBindProjectPath, "./my-path"))
		assert.NoError(t, os.Setenv(EnvDisableDocker, "true"))
		assert.NoError(t, os.Setenv(EnvCustomRulesPath, "test"))
		assert.NoError(t, os.Setenv(EnvEnableInformationSeverity, "true"))
		assert.NoError(t, os.Setenv(EnvEnableOwaspDependencyCheck, "true"))
		assert.NoError(t, os.Setenv(EnvEnableShellCheck, "true"))
		assert.NoError(t, os.Setenv(EnvShowVulnerabilitiesTypes, fmt.Sprintf("%s, %s", vulnerability.Vulnerability.ToString(), vulnerability.RiskAccepted.ToString())))
		configs.NewConfigsFromEnvironments()
		assert.Equal(t, configFilePath, configs.GetConfigFilePath())
		assert.Equal(t, "http://horusec.com", configs.GetHorusecAPIUri())
		assert.Equal(t, int64(99), configs.GetTimeoutInSecondsRequest())
		assert.Equal(t, int64(999), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, int64(20), configs.GetMonitorRetryInSeconds())
		assert.Equal(t, authorization, configs.GetRepositoryAuthorization())
		assert.Equal(t, "sonarqube", configs.GetPrintOutputType())
		assert.Equal(t, "./output-sonarqube.json", configs.GetJSONOutputFilePath())
		assert.Equal(t, []string{"INFO"}, configs.GetSeveritiesToIgnore())
		assert.Equal(t, []string{"**/*_test.go", "**/*_mock.go"}, configs.GetFilesOrPathsToIgnore())
		assert.Equal(t, false, configs.GetReturnErrorIfFoundVulnerability())
		assert.Equal(t, "./horusec-manager", configs.GetProjectPath())
		assert.Equal(t, workdir.NewWorkDir(), configs.GetWorkDir())
		assert.Equal(t, false, configs.GetEnableGitHistoryAnalysis())
		assert.Equal(t, false, configs.GetCertInsecureSkipVerify())
		assert.Equal(t, "./", configs.GetCertPath())
		assert.Equal(t, false, configs.GetEnableCommitAuthor())
		assert.Equal(t, "my-project", configs.GetRepositoryName())
		assert.Equal(t, []string{"hash7", "hash6"}, configs.GetRiskAcceptHashes())
		assert.Equal(t, []string{"hash9", "hash8"}, configs.GetFalsePositiveHashes())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString(), vulnerability.RiskAccepted.ToString()}, configs.GetShowVulnerabilitiesTypes())
		assert.Equal(t, map[string]string{"x-auth": "987654321"}, configs.GetHeaders())
		assert.Equal(t, "./my-path", configs.GetContainerBindProjectPath())
		assert.Equal(t, true, configs.GetDisableDocker())
		assert.Equal(t, "test", configs.GetCustomRulesPath())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
		assert.Equal(t, true, configs.GetEnableOwaspDependencyCheck())
		assert.Equal(t, true, configs.GetEnableShellCheck())
		cobraCmd := &cobra.Command{
			Use:     "start",
			Short:   "Start horusec-cli",
			Long:    "Start the Horusec' analysis in the current path",
			Example: "horusec start",
			RunE: func(cmd *cobra.Command, args []string) error {
				return nil
			},
		}
		_ = cobraCmd.PersistentFlags().
			StringP("project-path", "p", configs.GetProjectPath(), "Path to run an analysis in your project")
		_ = cobraCmd.PersistentFlags().
			StringSliceP("false-positive", "F", configs.GetFalsePositiveHashes(), "Used to ignore a vulnerability by hash and setting it to be of the false positive type. Example -F=\"hash1, hash2\"")
		_ = cobraCmd.PersistentFlags().
			StringSliceP("risk-accept", "R", configs.GetRiskAcceptHashes(), "Used to ignore a vulnerability by hash and setting it to be of the risk accept type. Example -R=\"hash3, hash4\"")
		_ = cobraCmd.PersistentFlags().
			Int64P("analysis-timeout", "t", configs.GetTimeoutInSecondsAnalysis(), "The timeout threshold for the Horusec CLI wait for the analysis to complete.")
		_ = cobraCmd.PersistentFlags().
			BoolP("information-severity", "I", configs.GetEnableInformationSeverity(), "Used to enable or disable information severity vulnerabilities, information vulnerabilities can contain a lot of false positives. Example: -I=\"true\"")
		_ = cobraCmd.PersistentFlags().
			StringSliceP("show-vulnerabilities-types", "", configs.GetShowVulnerabilitiesTypes(), "Used to show in the output vulnerabilities of types: Vulnerability, Risk Accepted, False Positive, Corrected. Example --show-vulnerabilities-types=\"Vulnerability, Risk Accepted\"")
		args := []string{"-p", "/home/usr/project", "-F", "SOMEHASHALEATORY1,SOMEHASHALEATORY2", "-R", "SOMEHASHALEATORY3,SOMEHASHALEATORY4", "-t", "1000", "I", "true", "--show-vulnerabilities-types", "Vulnerability"}
		assert.NoError(t, cobraCmd.PersistentFlags().Parse(args))
		assert.NoError(t, cobraCmd.Execute())
		configs.NewConfigsFromCobraAndLoadsCmdStartFlags(cobraCmd)
		assert.Equal(t, "/home/usr/project", configs.GetProjectPath())
		assert.Equal(t, []string{"SOMEHASHALEATORY1", "SOMEHASHALEATORY2"}, configs.GetFalsePositiveHashes())
		assert.Equal(t, []string{"SOMEHASHALEATORY3", "SOMEHASHALEATORY4"}, configs.GetRiskAcceptHashes())
		assert.Equal(t, []string{vulnerability.Vulnerability.ToString()}, configs.GetShowVulnerabilitiesTypes())
		assert.Equal(t, int64(1000), configs.GetTimeoutInSecondsAnalysis())
		assert.Equal(t, true, configs.GetEnableInformationSeverity())
	})
}

func TestToLowerCamel(t *testing.T) {
	t.Run("should success set all configs as lower camel case", func(t *testing.T) {
		configs := &Config{}
		assert.Equal(t, "horusecCliHorusecApiUri", configs.toLowerCamel(EnvHorusecAPIUri))
		assert.Equal(t, "horusecCliRepositoryAuthorization", configs.toLowerCamel(EnvRepositoryAuthorization))
		assert.Equal(t, "horusecCliCertPath", configs.toLowerCamel(EnvCertPath))
		assert.Equal(t, "horusecCliRepositoryName", configs.toLowerCamel(EnvRepositoryName))
		assert.Equal(t, "horusecCliPrintOutputType", configs.toLowerCamel(EnvPrintOutputType))
		assert.Equal(t, "horusecCliJsonOutputFilepath", configs.toLowerCamel(EnvJSONOutputFilePath))
		assert.Equal(t, "horusecCliProjectPath", configs.toLowerCamel(EnvProjectPath))
		assert.Equal(t, "horusecCliCustomRulesPath", configs.toLowerCamel(EnvCustomRulesPath))
		assert.Equal(t, "horusecCliContainerBindProjectPath", configs.toLowerCamel(EnvContainerBindProjectPath))
		assert.Equal(t, "horusecCliTimeoutInSecondsRequest", configs.toLowerCamel(EnvTimeoutInSecondsRequest))
		assert.Equal(t, "horusecCliTimeoutInSecondsAnalysis", configs.toLowerCamel(EnvTimeoutInSecondsAnalysis))
		assert.Equal(t, "horusecCliMonitorRetryInSeconds", configs.toLowerCamel(EnvMonitorRetryInSeconds))
		assert.Equal(t, "horusecCliReturnErrorIfFoundVulnerability", configs.toLowerCamel(EnvReturnErrorIfFoundVulnerability))
		assert.Equal(t, "horusecCliEnableGitHistoryAnalysis", configs.toLowerCamel(EnvEnableGitHistoryAnalysis))
		assert.Equal(t, "horusecCliCertInsecureSkipVerify", configs.toLowerCamel(EnvCertInsecureSkipVerify))
		assert.Equal(t, "horusecCliEnableCommitAuthor", configs.toLowerCamel(EnvEnableCommitAuthor))
		assert.Equal(t, "horusecCliDisableDocker", configs.toLowerCamel(EnvDisableDocker))
		assert.Equal(t, "horusecCliEnableInformationSeverity", configs.toLowerCamel(EnvEnableInformationSeverity))
		assert.Equal(t, "horusecCliSeveritiesToIgnore", configs.toLowerCamel(EnvSeveritiesToIgnore))
		assert.Equal(t, "horusecCliFilesOrPathsToIgnore", configs.toLowerCamel(EnvFilesOrPathsToIgnore))
		assert.Equal(t, "horusecCliFalsePositiveHashes", configs.toLowerCamel(EnvFalsePositiveHashes))
		assert.Equal(t, "horusecCliRiskAcceptHashes", configs.toLowerCamel(EnvRiskAcceptHashes))
		assert.Equal(t, "horusecCliToolsConfig", configs.toLowerCamel(EnvToolsConfig))
		assert.Equal(t, "horusecCliHeaders", configs.toLowerCamel(EnvHeaders))
		assert.Equal(t, "horusecCliWorkDir", configs.toLowerCamel(EnvWorkDir))
		assert.Equal(t, "horusecCliCustomImages", configs.toLowerCamel(EnvCustomImages))
		assert.Equal(t, "horusecCliShowVulnerabilitiesTypes", configs.toLowerCamel(EnvShowVulnerabilitiesTypes))
		assert.Equal(t, "horusecCliEnableOwaspDependencyCheck", configs.toLowerCamel(EnvEnableOwaspDependencyCheck))
		assert.Equal(t, "horusecCliEnableShellcheck", configs.toLowerCamel(EnvEnableShellCheck))
	})
}

func TestNormalizeConfigs(t *testing.T) {
	t.Run("Should success normalize config", func(t *testing.T) {
		config := &Config{}
		config.SetJSONOutputFilePath("./cli")
		config.SetProjectPath("./cli")

		assert.NotEmpty(t, config.NormalizeConfigs())
	})
}

func TestConfig_ToBytes(t *testing.T) {
	t.Run("Should success when parse config to json bytes without indent", func(t *testing.T) {
		config := &Config{}
		config.NewConfigsFromEnvironments()
		assert.NotEmpty(t, config.ToBytes(false))
	})
	t.Run("Should success when parse config to json bytes with indent", func(t *testing.T) {
		config := &Config{}
		config.NewConfigsFromEnvironments()
		assert.NotEmpty(t, config.ToBytes(true))
	})
}
func TestSetLogOutput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("Should fail when log path is invalid", func(t *testing.T) {
		config := NewConfig()
		config.SetLogFilePath("invalidPath")

		err := config.SetLogOutput()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "no such file or directory")
	})
	t.Run("Should fail when log path is a file", func(t *testing.T) {
		config := NewConfig()
		file, err := os.Create("./test.txt")
		assert.NoError(t, err)
		config.SetLogFilePath(file.Name())

		err = config.SetLogOutput()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "not a directory")
		_ = os.Remove(file.Name())
	})

	t.Run("Should success when log path is empty", func(t *testing.T) {
		config := NewConfig()
		err := config.SetLogOutput()
		assert.NoError(t, err)
	})
	t.Run("Should success when log path is valid", func(t *testing.T) {
		config := NewConfig()
		config.SetLogFilePath("./")
		err := config.SetLogOutput()
		assert.NoError(t, err)
	})

	t.Run("Should fail when get working directory fails", func(t *testing.T) {
		config := NewConfig()
		sysCallMock := mock_config.NewMockISystemCalls(ctrl)
		config.SetSystemCall(sysCallMock)

		expetedError := errors.New("error getting working directory")
		sysCallMock.EXPECT().Getwd().Return("", expetedError)
		err := config.SetLogOutput()
		assert.Error(t, err)
		assert.Equal(t, expetedError, err)
	})
	t.Run("Should fail when make directory fails", func(t *testing.T) {
		config := NewConfig()
		sysCallMock := mock_config.NewMockISystemCalls(ctrl)
		config.SetSystemCall(sysCallMock)

		expetedError := errors.New("error making directory")
		sysCallMock.EXPECT().Getwd().Return("", nil)
		sysCallMock.EXPECT().Stat(gomock.Any()).Return(nil, nil)
		sysCallMock.EXPECT().IsNotExist(gomock.Any()).Return(true)
		sysCallMock.EXPECT().MkdirAll(gomock.Any(), gomock.Any()).Return(expetedError)
		err := config.SetLogOutput()
		assert.Error(t, err)
		assert.Equal(t, expetedError, err)
	})
	t.Run("Should fail when create file fails", func(t *testing.T) {
		config := NewConfig()
		sysCallMock := mock_config.NewMockISystemCalls(ctrl)
		config.SetSystemCall(sysCallMock)

		expetedError := errors.New("error creating file")
		sysCallMock.EXPECT().Getwd().Return("", nil)
		sysCallMock.EXPECT().Stat(gomock.Any()).Return(nil, nil)
		sysCallMock.EXPECT().IsNotExist(gomock.Any()).Return(false)
		sysCallMock.EXPECT().Create(gomock.Any()).Return(nil, expetedError)
		err := config.SetLogOutput()
		assert.Error(t, err)
		assert.Equal(t, expetedError, err)
	})

}
func TestSetLogPath(t *testing.T) {
	t.Run("Should success when log path is not empty", func(t *testing.T) {
		config := &Config{}
		config.SetLogFilePath("aa")
		assert.Equal(t, "aa", config.GetLogFilePath())
	})
}

func TestLogLevel(t *testing.T) {
	t.Run("Should success when log level is not empty", func(t *testing.T) {
		config := &Config{}
		config.SetLogLevel(logrus.WarnLevel.String())
		assert.Equal(t, logrus.WarnLevel.String(), config.GetLogLevel())
	})
	t.Run("Should get default when log level is empty", func(t *testing.T) {
		config := &Config{}
		assert.Equal(t, logrus.InfoLevel.String(), config.GetLogLevel())
	})
}
