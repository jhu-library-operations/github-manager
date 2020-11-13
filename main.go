//type (
// 	BillingInfo struct {
		
// 		AssetPacks int
// 		AvailableSeats int
// 		BandwidthQuota float32
// 		BandwidthUsage float32
// 		Seats int
// 		StorageQuota float32
// 		StorageUsage float32
// 		StorageUsagePercentage float32
// 		TotalAvailableLicenses int
// 		TotalLicenses int
// 	} `graphql:"... on EnterpriseBillingInfo"`

// 	OwnerInfoFragment struct {
// 		AffiliatedUsersWithTwoFactorDisabledExist bool
// 		AllowPrivateRepositoryForkingSetting string
// 		DefaultRepositoryPermissionSetting string
// 		IpAllowListEnabledSetting string
// 		IsUpdatingTwoFactorRequirement bool
// 		MembersCanChangeRepositoryVisibilitySetting string
// 		MembersCanCreateInternalRepositoriesSetting bool
// 		MembersCanCreatePrivateRepositoriesSetting bool
// 		MembersCanCreatePublicRepositoriesSetting bool
// 		MembersCanCreateRepositoriesSetting string
// 		MembersCanDeleteIssuesSetting string
// 		MembersCanDeleteRepositoriesSetting string
// 		MembersCanInviteCollaboratorsSetting string
// 		MembersCanMakePurchasesSetting string
// 		MembersCanViewDependencyInsightsSetting string
// 		MembersCanUpdateProtectedBranchesSetting string
// 		OrganizationProjectsSetting string
// 		RepositoryProjectsSetting string
// 		SamlIdentityProvider struct {
// 			DigestMethod string
// 			Id string
// 			IdpCertificate string
// 			RecoveryCodes []string
// 			SignatureMethod string
// 			SsoUrl string
// 		}
// 	}

// 	EnterpriseFragment struct {
// 		AvatarUrl string
// 	}
		
// )

// func fetchEnterprise(client *githubv4.Client, slug string) (string, error) {
// 	var q struct {
// 		Enterprise struct {
// 			AvatarUrl string
// 			// BillingInfoFragment `graphql:"... on EnterpriseBillingInfo"`
// 			BillingInfo
// 			CreatedAt string
// 			DatabaseId int
// 			Description string
// 			Id string
// 			Location string
// 			Name string
// 			// OwnerInfoFragment `graphql:"... on EnterpriseOwnerInfo"`
// 			ResourcePath string
// 			Slug string
// 			Url string
// 			ViewerIsAdmin bool
// 			WebsiteUrl string
// 		}`graphql:"enterprise(slug: $slug)"`
// 	}

// 	vars := map[string]interface{}{
// 		"slug": githubv4.String(slug),
// 	}

// 	err := client.Query(context.Background(), &q, vars)
// 	if err != nil {
// 		return "", err
// 	}
	
// 	printJSON(q)
// 	return q.Enterprise.AvatarUrl, nil
// }

// func fetchOrganization(client *githubv4.Client, login string) (string, error){
// 	var q struct {
// 		Organization struct {
// 			AvatarUrl string
// 			CreatedAt string
// 			DatabaseId int
// 			Description string
// 			Email string
// 			Id string
// 			IpAllowListEnabledSetting string
// 			IsVerified bool
// 			ItemShowcase struct {
// 				HasPinnedItems bool
// 			}
// 			Location string
// 			Login string
// 			Name string
// 			NewTeamResourcePath string
// 			NewTeamUrl string
// 			OrganizationBillingEmail string
// 			PinnedItemsRemaining int
// 			// Project struct {
// 			// 	Columns struct {}
// 			// 	PendingCards struct {}
// 			// 	Body string
// 			// 	Closed bool
// 			// 	ClosedAt string
// 			// 	CreatedAt string
// 			// 	Creator struct {}
// 			// 	DatabaseId int
// 			// 	Id string
// 			// 	Name string
// 			// 	Number int
// 			// 	Owner struct {}
// 			// 	ResourcePath string
// 			// 	State struct {}
// 			// 	UpdatedAt string
// 			// 	Url string
// 			// 	ViewerCanUpdate bool
// 			// }
		
// 		}`graphql:"organization(login: $login)"`
// 	}

// 	vars := map[string]interface{}{
// 		"login": githubv4.String(login),
// 	}

// 	err := client.Query(context.Background(), &q, vars)
// 	if err != nil {
// 		return "", err
// 	}
	
// 	printJSON(q)
// 	return q.Organization.Name, nil	
// }

// type CreateEnterpriseOrgInput struct {
// 	EnterpriseId githubv4.String
// 	Login githubv4.String
// 	BillingEmail githubv4.String
// 	AdminLogins []githubv4.String
// }

// func printJSON(v interface{}) {
// 	w := json.NewEncoder(os.Stdout)
// 	w.SetIndent("", "\t")
// 	err := w.Encode(v)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func createEnterpriseOrg(client *githubv4.Client, input *CreateEnterpriseOrgInput) (string, error) {
// 	var m struct {
// 		CreateEnterpriseOrganization struct {
// 			ClientMutationId string
// 			Enterprise struct {
// 				Id string
// 			}	
// 			Organization struct {
// 				Id string
// 				Login string
// 				Name string
// 				Email string
// 			}
// 		}`graphql:"createEnterpriseOrganization(input: $input)"`
// 	}

// 	orginput := githubv4.CreateEnterpriseOrganizationInput{
// 		EnterpriseID: input.EnterpriseId,
// 		Login: input.Login,
// 		ProfileName: input.Login,
// 		BillingEmail: input.BillingEmail,
// 		AdminLogins: input.AdminLogins,
// 	}

// 	err := client.Mutate(context.Background(), &m, orginput, nil)
// 	if err != nil {
// 		return "", err
// 	}

// 	printJSON(m)
// 	return m.CreateEnterpriseOrganization.Organization.Login, nil
// }

// func main() {
// 	var bearer = os.Getenv("GH_TOKEN")
// 	src := oauth2.StaticTokenSource( &oauth2.Token{AccessToken: bearer })
// 	httpClient := oauth2.NewClient(context.Background(), src)
// 	client := githubv4.NewClient(httpClient)
	
// 	avatarUrl, err := fetchEnterprise(client, "johns-hopkins-university")
// 	if err != nil { panic(err) }
// 	fmt.Println(avatarUrl)

// 	name, err := fetchOrganization(client, "jhu-library-operations")
// 	if err != nil { panic(err) }
// 	fmt.Println(name)
// 	// login, err := fetchOrganization(client, "juh-sheridan-libraries")
// 	// fmt.Println(login)

// 	// input := CreateEnterpriseOrgInput{
// 	// 	EnterpriseId: githubv4.String(enterpriseId),
// 	// 	Login: "juh-library-test",
// 	// 	BillingEmail: "dbelrose@jhu.edu",
// 	// }
// 	// input.AdminLogins = append(input.AdminLogins, "derekbelrose")
// 	// if err != nil {
// 	// 	resp, err := createEnterpriseOrg(client, &input)
// 	// 	if err != nil {
// 	// 		panic(err)
// 	// 	}
// 	// 	fmt.Println(resp)
// 	// 	fmt.Println("Organization created, Yay!")
// 	// }


	
// }

package main 
import "github-manager/cmd"

func main() {
	cmd.Execute()
}
