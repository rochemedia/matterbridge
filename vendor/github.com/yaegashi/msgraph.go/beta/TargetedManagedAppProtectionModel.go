// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TargetedManagedAppProtection Policy used to configure detailed management settings targeted to specific security groups
type TargetedManagedAppProtection struct {
	// ManagedAppProtection is the base model of TargetedManagedAppProtection
	ManagedAppProtection
	// IsAssigned Indicates if the policy is deployed to any inclusion groups or not.
	IsAssigned *bool `json:"isAssigned,omitempty"`
	// TargetedAppManagementLevels The intended app management levels for this policy
	TargetedAppManagementLevels *AppManagementLevel `json:"targetedAppManagementLevels,omitempty"`
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}