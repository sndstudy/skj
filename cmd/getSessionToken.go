/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"unicode/utf8"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sndstudy/skj/pkg/awswrap"
	"github.com/spf13/cobra"
)

// getSessionTokenCmd represents the getSessionToken command
var getSessionTokenCmd = &cobra.Command{
	Use:   "get-session-token",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		profileName, _ := cmd.Flags().GetString("profile")
		serialNumber, _ := cmd.Flags().GetString("serial-number")
		tokenCode, _ := cmd.Flags().GetString("token-code")
		durationSeconds, _ := cmd.Flags().GetInt32("duration-seconds")

		input := &sts.GetSessionTokenInput{
			DurationSeconds: aws.Int32(durationSeconds),
		}

		if utf8.RuneCountInString(serialNumber) != 0 {
			input.SerialNumber = aws.String(serialNumber)
		}

		if utf8.RuneCountInString(tokenCode) != 0 {
			input.TokenCode = aws.String(tokenCode)
		}

		client, err := awswrap.GetStsClient(profileName)
		if err != nil {
			return err
		}

		output, err := awswrap.GetSessionToken(context.Background(), client, input)
		if err != nil {
			return err
		}
		log.Println(*output.Credentials.AccessKeyId, *output.Credentials.SecretAccessKey, *output.Credentials.SessionToken)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getSessionTokenCmd)

	getSessionTokenCmd.Flags().StringP("profile", "", "default", "Use a specific profile from your credential file.")
	getSessionTokenCmd.Flags().StringP("serial-number", "", "", "The identification number of the MFA device that is associated with the IAM user who is making the GetSessionToken call.")
	getSessionTokenCmd.Flags().StringP("token-code", "", "", "The value provided by the MFA device, if MFA is required.")
	getSessionTokenCmd.Flags().Int32P("duration-seconds", "", 43200, "The  duration, in seconds, that the credentials should remain valid.")
}
