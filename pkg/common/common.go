package common

import "github.com/spf13/viper"

// ReadViperConfigFile is define the config paths and read the configuration file.
func ReadViperConfigFile(configName string) {
	viper.SetConfigName(configName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("/tk8")
	verr := viper.ReadInConfig() // Find and read the config file.
	ErrorCheck("no config provided", verr)
}

// GetCredentials get the aws credentials from the config file.
func GetCredentials() AwsCredentials {
	ReadViperConfigFile("config")
	return AwsCredentials{
		AwsAccessKeyID:   viper.GetString("aws.aws_access_key_id"),
		AwsSecretKey:     viper.GetString("aws.aws_secret_access_key"),
		AwsAccessSSHKey:  viper.GetString("aws.aws_ssh_keypair"),
		AwsDefaultRegion: viper.GetString("aws.aws_default_region"),
	}
}
