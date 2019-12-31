package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func StubProfiles(startIndex int, endIndex int) []AWSProfile {
	var profiles []AWSProfile

	for i := startIndex; i <= endIndex; i++ {
		profiles = append(profiles, AWSProfile{
			ProfileName:        fmt.Sprintf("profile-%d", i),
			DisplayProfileName: fmt.Sprintf("profile-%d display name", i),
		})
	}

	return profiles
}

func TestFindProfileInCredentialsFile_ReturnsNilIfNoProfileWithGivenNameFound(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles:   StubProfiles(1, 2),
		ConfigAssumedProfiles: []AWSProfile{},
	}

	result := profiles.FindProfileInCredentialsFile("profile-3")

	assert.Nil(t, result)
}

func TestFindProfileInCredentialsFile_ReturnsProfileOnceFound(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles:   StubProfiles(1, 2),
		ConfigAssumedProfiles: []AWSProfile{},
	}

	result := profiles.FindProfileInCredentialsFile("profile-2")

	assert.NotNil(t, result)
	assert.Equal(t, result.ProfileName, "profile-2")
}

func TestFindProfileInConfigFile_ReturnsNilIfNoProfileWithGivenNameFound(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles:   []AWSProfile{},
		ConfigAssumedProfiles: StubProfiles(1, 2),
	}

	result := profiles.FindProfileInConfigFile("profile-3")

	assert.Nil(t, result)
}

func TestFindProfileInConfigFile_ReturnsProfileOnceFound(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles:   []AWSProfile{},
		ConfigAssumedProfiles: StubProfiles(1, 2),
	}

	result := profiles.FindProfileInConfigFile("profile-2")

	assert.NotNil(t, result)
	assert.Equal(t, result.ProfileName, "profile-2")
}

func TestGetAllDisplayProfileNames_ReturnsProfileNamesFromBothCredentialsAndConfigFiles(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles:   StubProfiles(1, 2),
		ConfigAssumedProfiles: StubProfiles(3, 3),
	}

	result := profiles.GetAllDisplayProfileNames()

	expected := []string{
		"profile-1 display name",
		"profile-2 display name",
		"profile-3 display name",
	}
	assert.ElementsMatch(t, expected, result)
}

func TestFilter_ReturnsFilteredProfilesFromBothCredentialsAndConfigFiles(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles: []AWSProfile{
			{
				ProfileName:        "credentials profile 1",
				DisplayProfileName: "credentials profile 1",
			},
			{
				ProfileName:        "credentials profile 2 - match",
				DisplayProfileName: "credentials profile 2",
			},
		},
		ConfigAssumedProfiles: []AWSProfile{
			{
				ProfileName:        "config profile 1",
				DisplayProfileName: "config profile 1",
			},
			{
				ProfileName:        "config profile 2 - match",
				DisplayProfileName: "config profile 2",
			},
		},
	}

	result := profiles.Filter("match")

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "credentials profile 2 - match", result[0].ProfileName)
	assert.Equal(t, "config profile 2 - match", result[1].ProfileName)
}

func TestFilter_ReturnsEmptyIfNoneMatchesFilter(t *testing.T) {
	profiles := AWSProfiles{
		CredentialsProfiles: []AWSProfile{
			{
				ProfileName:        "credentials profile 1",
				DisplayProfileName: "credentials profile 1",
			},
			{
				ProfileName:        "credentials profile 2",
				DisplayProfileName: "credentials profile 2",
			},
		},
		ConfigAssumedProfiles: []AWSProfile{
			{
				ProfileName:        "config profile 1",
				DisplayProfileName: "config profile 1",
			},
			{
				ProfileName:        "config profile 2",
				DisplayProfileName: "config profile 2",
			},
		},
	}

	result := profiles.Filter("match")

	assert.Equal(t, 0, len(result))
}
