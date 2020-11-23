package parse

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/resourceid"
)

var _ resourceid.Formatter = TimeSeriesInsightsAccessPolicyId{}

func TestTimeSeriesInsightsAccessPolicyIDFormatter(t *testing.T) {
	subscriptionId := "12345678-1234-5678-1234-123456789012"
	actual := NewTimeSeriesInsightsAccessPolicyID(subscriptionId, "resourceGroup1", "env1", "policy1").ID("")
	expected := "/subscriptions/12345678-1234-5678-1234-123456789012/resourceGroups/resourceGroup1/Microsoft.TimeSeriesInsights/environments/env1/accessPolicies/policy1"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestTimeSeriesInsightsAccessPolicyId(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *TimeSeriesInsightsAccessPolicyId
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Segment",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/",
			Expected: nil,
		},
		{
			Name:     "Resource Group ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo/",
			Expected: nil,
		},
		{
			Name:     "Time Series Insight AccessPolicy Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.TimeSeriesInsights/environments/Environment1/accessPolicies/",
			Expected: nil,
		},
		{
			Name:     "Time Series Insight Access Policy ID legacy casing",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.TimeSeriesInsights/environments/Environment1/accesspolicies/Policy1",
			Expected: nil,
		},
		{
			Name:  "Time Series Insight Access Policy ID",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.TimeSeriesInsights/environments/Environment1/accessPolicies/Policy1",
			Expected: &TimeSeriesInsightsAccessPolicyId{
				SubscriptionId:  "00000000-0000-0000-0000-000000000000",
				EnvironmentName: "Environment1",
				ResourceGroup:   "resGroup1",
				Name:            "Policy1",
			},
		},
		{
			Name:     "Wrong Casing",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.TimeSeriesInsights/Environments/Environment1/AccessPolicies/Policy1",
			Expected: nil,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := TimeSeriesInsightsAccessPolicyID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for Resource Group", v.Expected.ResourceGroup, actual.ResourceGroup)
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for Subscription Id", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
	}
}
