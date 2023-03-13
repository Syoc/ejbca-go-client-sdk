package main

type ConfigDumpModel struct {
	EndEntityProfiles struct {
		TlsServerCompany1 struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerCompany1"`
		HaydenEndEntity struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				UNIFORMRESOURCEID []struct {
					Required bool `json:"Required"`
				} `json:"UNIFORMRESOURCEID"`
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"haydenEndEntity"`
		MsaeWorkstations struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"msaeWorkstations"`
		TlsServerIT struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerIT"`
		TlsServerCompany2 struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerCompany2"`
		ClevelandWorkstation2D struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailRequired                                  bool     `json:"E-Mail Required"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			Notifications                                  []struct {
				Subject   string   `json:"Subject"`
				Events    []string `json:"Events"`
				Sender    string   `json:"Sender"`
				Recipient string   `json:"Recipient"`
				Message   string   `json:"Message"`
			} `json:"Notifications"`
			SubjectDN struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"cleveland-workstation2d"`
		SolitudeDomainController struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value,omitempty"`
				} `json:"DC"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				GUID []struct {
					Required bool `json:"Required"`
				} `json:"GUID"`
			} `json:"Subject Alternative Name"`
		} `json:"solitude-DomainController"`
		CarlCMP struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"OU"`
				O []struct {
					Required bool `json:"Required"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				RFC822NAME []struct {
					Required bool `json:"Required"`
					Use      bool `json:"Use"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"carlCMP"`
		UserAuthenticationApproval struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				UPN []struct {
					Required bool `json:"Required"`
				} `json:"UPN"`
			} `json:"Subject Alternative Name"`
		} `json:"userAuthenticationApproval"`
		DomainController struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"domainController"`
		TimeStampingAuthority struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
		} `json:"timeStampingAuthority"`
		ScepDevice struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"scepDevice"`
		TlsServerAnyCA struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerAnyCA"`
		UserAuthenticationAzure struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				RFC822NAME []struct {
					Required bool `json:"Required"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"userAuthenticationAzure"`
		RemoteDesktopAuth struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"remoteDesktopAuth"`
		TestDeleteMe struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			AllowMergeDN                                   bool     `json:"Allow Merge DN"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable,omitempty"`
					Values     []string `json:"Values,omitempty"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"test-delete-me"`
		MsaeUserAuthentication struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required bool `json:"Required"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				UPN []struct {
					Required bool `json:"Required"`
				} `json:"UPN"`
			} `json:"Subject Alternative Name"`
		} `json:"msaeUserAuthentication"`
		IotDevice struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				O []struct {
					Required bool `json:"Required"`
				} `json:"O"`
				L []struct {
					Required bool `json:"Required"`
				} `json:"L"`
				ST []struct {
					Required bool `json:"Required"`
				} `json:"ST"`
				C []struct {
					Required bool `json:"Required"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				RFC822NAME []struct {
					Required bool `json:"Required"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"iot_device"`
		TlsServerMgmt struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				O []struct {
					Required bool `json:"Required"`
				} `json:"O"`
				C []struct {
					Required bool `json:"Required"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				IPADDRESS []struct {
					Required bool `json:"Required"`
				} `json:"IPADDRESS"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerMgmt"`
		VpnFirewallAutoUname10M struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			UsernameAutoGenerated                          bool     `json:"Username Auto-Generated"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				IPADDRESS []struct {
					Required bool `json:"Required"`
				} `json:"IPADDRESS"`
			} `json:"Subject Alternative Name"`
		} `json:"vpnFirewallAutoUname10m"`
		CodeSigner struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"O"`
				OU []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"OU"`
				UID []struct {
					Required bool `json:"Required"`
				} `json:"UID"`
			} `json:"Subject DN"`
		} `json:"codeSigner"`
		VpnFirewall10M struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				IPADDRESS []struct {
					Required bool `json:"Required"`
				} `json:"IPADDRESS"`
			} `json:"Subject Alternative Name"`
		} `json:"vpnFirewall10m"`
		TestMultiIssue struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			CleartextPasswordEnabled     bool     `json:"Cleartext Password Enabled"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool `json:"Required"`
					Validation struct {
						Regex string `json:"Regex"`
					} `json:"Validation,omitempty"`
				} `json:"OU"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"testMultiIssue"`
		APasswordTest struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"aPasswordTest"`
		MsaeDomainController struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value,omitempty"`
				} `json:"DC"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value,omitempty"`
				} `json:"DNSNAME"`
				GUID []struct {
					Required bool `json:"Required"`
				} `json:"GUID"`
			} `json:"Subject Alternative Name"`
		} `json:"msaeDomainController"`
		EstDevice struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"estDevice"`
		TlsServerApprovalIT struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerApprovalIT"`
		UserEncryption2 struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			KeyRecoveryUsed              bool     `json:"Key Recovery Used"`
			KeyRecoveryRequired          bool     `json:"Key Recovery Required"`
			ReuseKeyRecoveredCertificate bool     `json:"Reuse Key-Recovered Certificate"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				RFC822NAME []struct {
					Required bool `json:"Required"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"userEncryption2"`
		IkeIntermed struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"ike-intermed"`
		TlsPeerConnMgmt struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				IPADDRESS []struct {
					Required bool `json:"Required"`
				} `json:"IPADDRESS"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsPeerConnMgmt"`
		MsaeTlsApproval struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			AllowMergeDN                                   bool     `json:"Allow Merge DN"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"msaeTlsApproval"`
		MsaeWorkstationsEFS struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			KeyRecoveryUsed                                bool     `json:"Key Recovery Used"`
			KeyRecoveryDefault                             bool     `json:"Key Recovery Default"`
			KeyRecoveryRequired                            bool     `json:"Key Recovery Required"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				RFC822NAME []struct {
					Required bool `json:"Required"`
					Use      bool `json:"Use"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"msaeWorkstationsEFS"`
		ExternalCA struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
		} `json:"External-CA"`
		VpnFirewallTest1D struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"O"`
				OU []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				IPADDRESS []struct {
					Required bool `json:"Required"`
				} `json:"IPADDRESS"`
			} `json:"Subject Alternative Name"`
		} `json:"vpnFirewallTest1d"`
		SignersIT struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
		} `json:"signersIT"`
		TlsServerAcme60D struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			UsernameAutoGenerated                          bool     `json:"Username Auto-Generated"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerAcme60d"`
		TlsServerAcme struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			UsernameAutoGenerated                          bool     `json:"Username Auto-Generated"`
			PasswordRequired                               bool     `json:"Password Required"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerAcme"`
		TlsServerEC struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServer-EC"`
		SolitudeWorkstationsMsae struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value,omitempty"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"solitude-workstations-msae"`
		UserAuthentication struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				UPN []struct {
					Required bool `json:"Required"`
				} `json:"UPN"`
			} `json:"Subject Alternative Name"`
		} `json:"userAuthentication"`
		MsaeUserDigitalSignature struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				RFC822NAME []struct {
					Required bool `json:"Required"`
					Use      bool `json:"Use"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"msaeUserDigitalSignature"`
		SolnaTlsServerAcme2D struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"solna-tlsServerAcme2d"`
		Keyrecov struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			KeyRecoveryUsed                                bool     `json:"Key Recovery Used"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"keyrecov"`
		TreasuryEmailValidation struct {
			ObjectType                   string        `json:"Object Type"`
			Version                      int           `json:"Version"`
			Name                         string        `json:"Name"`
			EMailRequired                bool          `json:"E-Mail Required"`
			EMailModifiable              bool          `json:"E-Mail Modifiable"`
			EMailDomain                  string        `json:"E-Mail Domain"`
			Description                  string        `json:"Description"`
			DefaultCA                    string        `json:"Default CA"`
			AvailableCAs                 []interface{} `json:"Available CAs"`
			DefaultCertificateProfile    string        `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string      `json:"Available Certificate Profiles"`
			DefaultTokenType             string        `json:"Default Token Type"`
			AvailableTokenTypes          []string      `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
			} `json:"Subject DN"`
		} `json:"treasuryEmailValidation"`
		IotCa struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				O []struct {
					Required bool `json:"Required"`
				} `json:"O"`
				L []struct {
					Required bool `json:"Required"`
				} `json:"L"`
				ST []struct {
					Required bool `json:"Required"`
				} `json:"ST"`
				C []struct {
					Required bool `json:"Required"`
				} `json:"C"`
			} `json:"Subject DN"`
		} `json:"iot_ca"`
		OcspSigner struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"ocspSigner"`
		AdminInternal struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
		} `json:"AdminInternal"`
		IntuneDevice struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"intuneDevice"`
		XcertCa struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			UseExtensionData             bool     `json:"Use Extension Data"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"OU"`
				O []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"O"`
				C []struct {
					Required     bool   `json:"Required"`
					DefaultValue string `json:"Default Value"`
				} `json:"C"`
			} `json:"Subject DN"`
		} `json:"xcert_ca"`
		UserEncryption struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			KeyRecoveryUsed              bool     `json:"Key Recovery Used"`
			KeyRecoveryDefault           bool     `json:"Key Recovery Default"`
			KeyRecoveryRequired          bool     `json:"Key Recovery Required"`
			ReuseKeyRecoveredCertificate bool     `json:"Reuse Key-Recovered Certificate"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				RFC822NAME []struct {
					Required bool `json:"Required"`
					Use      bool `json:"Use,omitempty"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"userEncryption"`
		SolitudeUsersMsae struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required bool `json:"Required"`
				} `json:"OU"`
				DC []struct {
					Required bool `json:"Required"`
				} `json:"DC"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				UPN []struct {
					Required bool `json:"Required"`
				} `json:"UPN"`
			} `json:"Subject Alternative Name"`
		} `json:"solitude-users-msae"`
		VpnFirewall struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
				IPADDRESS []struct {
					Required bool `json:"Required"`
				} `json:"IPADDRESS"`
			} `json:"Subject Alternative Name"`
		} `json:"vpnFirewall"`
		SshHost struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
			} `json:"Subject DN"`
			SshPrincipals struct {
				PRINCIPAL []struct {
					Required bool `json:"Required"`
				} `json:"PRINCIPAL"`
			} `json:"Ssh Principals"`
		} `json:"sshHost"`
		UserDigitalSignature struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			EMailUsed                    bool     `json:"E-Mail Used"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				RFC822NAME []struct {
					Required bool `json:"Required"`
				} `json:"RFC822NAME"`
			} `json:"Subject Alternative Name"`
		} `json:"userDigitalSignature"`
		XCertOrg1 struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []string `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
			} `json:"Subject DN"`
		} `json:"xCert-Org1"`
		TlsServerAcme2D struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			UsernameAutoGenerated                          bool     `json:"Username Auto-Generated"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerAcme2d"`
		SshUser struct {
			ObjectType                   string   `json:"Object Type"`
			Version                      int      `json:"Version"`
			Name                         string   `json:"Name"`
			UsernameAutoGenerated        bool     `json:"Username Auto-Generated"`
			PasswordAutoGenerated        bool     `json:"Password Auto-Generated"`
			PasswordRequired             bool     `json:"Password Required"`
			Description                  string   `json:"Description"`
			DefaultCA                    string   `json:"Default CA"`
			AvailableCAs                 []string `json:"Available CAs"`
			DefaultCertificateProfile    int      `json:"Default Certificate Profile"`
			AvailableCertificateProfiles []int    `json:"Available Certificate Profiles"`
			DefaultTokenType             string   `json:"Default Token Type"`
			AvailableTokenTypes          []string `json:"Available Token Types"`
			SubjectDN                    struct {
			} `json:"Subject DN"`
			SshPrincipals struct {
				PRINCIPAL []struct {
					Required bool `json:"Required"`
				} `json:"PRINCIPAL"`
			} `json:"Ssh Principals"`
		} `json:"sshUser"`
		Workstations struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			Notifications                                  []struct {
				Subject   string   `json:"Subject"`
				Events    []string `json:"Events"`
				Sender    string   `json:"Sender"`
				Recipient string   `json:"Recipient"`
				Message   string   `json:"Message"`
			} `json:"Notifications"`
			SubjectDN struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"workstations"`
		TlsServerDropAtt struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			UsernameAutoGenerated                          bool     `json:"Username Auto-Generated"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerDropAtt"`
		SolnaWorkstation struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailRequired                                  bool     `json:"E-Mail Required"`
			EMailDomain                                    string   `json:"E-Mail Domain"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			NotificationsUsed                              bool     `json:"Notifications Used"`
			Notifications                                  []struct {
				Subject   string   `json:"Subject"`
				Events    []string `json:"Events"`
				Sender    string   `json:"Sender"`
				Recipient string   `json:"Recipient"`
				Message   string   `json:"Message"`
			} `json:"Notifications"`
			SubjectDN struct {
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"solna-workstation"`
		AllowExtensionOverride struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			EMailUsed                                      bool     `json:"E-Mail Used"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				OU []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"OU"`
				O []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"O"`
				C []struct {
					Required   bool     `json:"Required"`
					Modifiable bool     `json:"Modifiable"`
					Values     []string `json:"Values"`
				} `json:"C"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
		} `json:"allowExtensionOverride"`
		TlsServerAcme30D struct {
			ObjectType                                     string   `json:"Object Type"`
			Version                                        int      `json:"Version"`
			Name                                           string   `json:"Name"`
			UsernameAutoGenerated                          bool     `json:"Username Auto-Generated"`
			CleartextPasswordEnabled                       bool     `json:"Cleartext Password Enabled"`
			MaxFailedLoginsModifiable                      bool     `json:"Max Failed Logins Modifiable"`
			Description                                    string   `json:"Description"`
			DefaultCA                                      string   `json:"Default CA"`
			AvailableCAs                                   []string `json:"Available CAs"`
			DefaultCertificateProfile                      string   `json:"Default Certificate Profile"`
			AvailableCertificateProfiles                   []string `json:"Available Certificate Profiles"`
			DefaultTokenType                               string   `json:"Default Token Type"`
			AvailableTokenTypes                            []string `json:"Available Token Types"`
			IssuanceRevocationReasonModifiable             bool     `json:"Issuance Revocation Reason Modifiable"`
			CABrowserForumOrganizationIdentifierModifiable bool     `json:"CA Browser Forum Organization Identifier Modifiable"`
			SubjectDN                                      struct {
				CN []struct {
					Required bool `json:"Required"`
				} `json:"CN"`
				E []struct {
					Required bool `json:"Required"`
				} `json:"E"`
			} `json:"Subject DN"`
			SubjectAlternativeName struct {
				DNSNAME []struct {
					Required bool `json:"Required"`
				} `json:"DNSNAME"`
			} `json:"Subject Alternative Name"`
		} `json:"tlsServerAcme30d"`
	} `json:"end-entity-profiles"`
	CustomCertificateExtensions struct {
		CertificateTemplateOID struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			ExtensionClass   string `json:"Extension class"`
			ObjectIdentifier string `json:"Object Identifier"`
			Critical         bool   `json:"Critical"`
			Required         bool   `json:"Required"`
			Properties       struct {
				Encoding string `json:"encoding"`
				Value    string `json:"value"`
				Dynamic  string `json:"dynamic"`
			} `json:"Properties"`
		} `json:"certificateTemplateOID"`
		PolicyMapping struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			ExtensionClass   string `json:"Extension class"`
			ObjectIdentifier string `json:"Object Identifier"`
			Critical         bool   `json:"Critical"`
			Required         bool   `json:"Required"`
			Properties       struct {
				Encoding string `json:"encoding"`
				Value    string `json:"value"`
				Dynamic  string `json:"dynamic"`
			} `json:"Properties"`
		} `json:"Policy Mapping"`
		CertificateTemplateInformation struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			ExtensionClass   string `json:"Extension class"`
			ObjectIdentifier string `json:"Object Identifier"`
			Critical         bool   `json:"Critical"`
			Required         bool   `json:"Required"`
			Properties       struct {
				Encoding string `json:"encoding"`
				Value    string `json:"value"`
				Dynamic  string `json:"dynamic"`
			} `json:"Properties"`
		} `json:"Certificate Template Information"`
		SzOIDNTDSCASECURITYEXT struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			ExtensionClass   string `json:"Extension class"`
			ObjectIdentifier string `json:"Object Identifier"`
			Critical         bool   `json:"Critical"`
			Required         bool   `json:"Required"`
			Properties       struct {
				Encoding string `json:"encoding"`
				Value    string `json:"value"`
				Dynamic  string `json:"dynamic"`
			} `json:"Properties"`
		} `json:"szOID_NTDS_CA_SECURITY_EXT"`
	} `json:"custom-certificate-extensions"`
	ApprovalProfiles struct {
		TlsCertificateApproval struct {
			ObjectType                      string `json:"Object Type"`
			Version                         int    `json:"Version"`
			ID                              int    `json:"ID"`
			Type                            string `json:"Type"`
			Name                            string `json:"Name"`
			ApprovalExpirationPeriod        string `json:"Approval Expiration Period"`
			RequestExpirationPeriod         string `json:"Request Expiration Period"`
			MaxExtensionTime                string `json:"Max Extension Time"`
			AllowSelfApprovedRequestEditing bool   `json:"Allow Self-Approved Request Editing"`
			NumberOfApprovalsRequired       int    `json:"Number of Approvals Required"`
		} `json:"tlsCertificateApproval"`
		TlsMgmtApproval struct {
			ObjectType                      string `json:"Object Type"`
			Version                         int    `json:"Version"`
			ID                              int    `json:"ID"`
			Type                            string `json:"Type"`
			Name                            string `json:"Name"`
			ApprovalExpirationPeriod        string `json:"Approval Expiration Period"`
			RequestExpirationPeriod         string `json:"Request Expiration Period"`
			MaxExtensionTime                string `json:"Max Extension Time"`
			AllowSelfApprovedRequestEditing bool   `json:"Allow Self-Approved Request Editing"`
			Steps                           []struct {
				StepIdentifier int `json:"Step Identifier"`
				Partitions     []struct {
					PartitionIdentifier int           `json:"Partition Identifier"`
					Title               string        `json:"Title"`
					CanApprove          []string      `json:"Can Approve"`
					CanView             []string      `json:"Can View"`
					Components          []interface{} `json:"Components"`
				} `json:"Partitions"`
			} `json:"Steps"`
		} `json:"tlsMgmtApproval"`
	} `json:"approval-profiles"`
	ExtendedKeyUsage struct {
		ExtendedKeyUsages struct {
			ObjectType           string `json:"Object Type"`
			Version              int    `json:"Version"`
			AllExtendedKeyUsages struct {
				Field1  string `json:"2.5.29.37.0"`
				Field2  string `json:"1.3.6.1.5.5.7.3.1"`
				Field3  string `json:"1.3.6.1.5.5.7.3.2"`
				Field4  string `json:"1.3.6.1.5.5.7.3.3"`
				Field5  string `json:"1.3.6.1.5.5.7.3.4"`
				Field6  string `json:"1.3.6.1.5.5.7.3.8"`
				Field7  string `json:"1.3.6.1.5.5.7.3.9"`
				Field8  string `json:"1.3.6.1.5.5.7.3.13"`
				Field9  string `json:"1.3.6.1.5.5.7.3.14"`
				Field10 string `json:"1.3.6.1.5.5.7.3.15"`
				Field11 string `json:"1.3.6.1.5.5.7.3.16"`
				Field12 string `json:"1.3.6.1.5.5.7.3.17"`
				Field13 string `json:"1.3.6.1.5.5.7.3.20"`
				Field14 string `json:"1.3.6.1.5.5.7.3.21"`
				Field15 string `json:"1.3.6.1.5.5.7.3.22"`
				Field16 string `json:"1.3.6.1.4.1.311.20.2.2"`
				Field17 string `json:"1.3.6.1.4.1.311.10.3.12"`
				Field18 string `json:"1.3.6.1.4.1.311.2.1.21"`
				Field19 string `json:"1.3.6.1.4.1.311.2.1.22"`
				Field20 string `json:"1.3.6.1.4.1.311.10.3.4"`
				Field21 string `json:"1.3.6.1.4.1.311.10.3.4.1"`
				Field22 string `json:"2.16.840.1.113741.1.2.3"`
				Field23 string `json:"0.4.0.2231.3.0"`
				Field24 string `json:"1.2.840.113583.1.1.5"`
				Field25 string `json:"1.2.203.7064.1.1.369791.1"`
				Field26 string `json:"1.2.203.7064.1.1.369791.2"`
				Field27 string `json:"1.3.6.1.5.2.3.4"`
				Field28 string `json:"1.3.6.1.5.2.3.5"`
				Field29 string `json:"2.23.136.1.1.3"`
				Field30 string `json:"2.23.136.1.1.8"`
				Field31 string `json:"2.16.840.1.101.3.6.8"`
				Field32 string `json:"1.3.6.1.4.74.4.26"`
				Field33 string `json:"1.3.6.1.5.5.8.2.2"`
			} `json:"All extended key usages"`
		} `json:"extended-key-usages"`
	} `json:"extended-key-usage"`
	Validators struct {
		BlocklistValidator struct {
			ObjectType                     string   `json:"Object Type"`
			Version                        int      `json:"Version"`
			Type                           string   `json:"Type"`
			Name                           string   `json:"Name"`
			Description                    string   `json:"Description"`
			IfValidationFails              string   `json:"If Validation Fails"`
			IfValidatorNotApplicable       string   `json:"If Validator Not Applicable"`
			IssuancePhase                  string   `json:"Issuance Phase"`
			ApplyForAllCertificateProfiles bool     `json:"Apply for All Certificate Profiles"`
			KeyAlgorithms                  []string `json:"Key Algorithms"`
		} `json:"Blocklist-Validator"`
		EcValidator struct {
			ObjectType                   string        `json:"Object Type"`
			Version                      int           `json:"Version"`
			Type                         string        `json:"Type"`
			Name                         string        `json:"Name"`
			Description                  string        `json:"Description"`
			IfValidationFails            string        `json:"If Validation Fails"`
			IfValidatorNotApplicable     string        `json:"If Validator Not Applicable"`
			IssuancePhase                string        `json:"Issuance Phase"`
			ApplyForCertificateProfiles  []interface{} `json:"Apply for Certificate Profiles"`
			KeyValidatorSettingsTemplate string        `json:"Key Validator Settings Template"`
			Curves                       []string      `json:"Curves"`
			UseFullPublicKeyValidation   bool          `json:"Use Full Public Key Validation"`
		} `json:"EcValidator"`
		AcmeAllowedDomains struct {
			ObjectType                     string        `json:"Object Type"`
			Version                        int           `json:"Version"`
			Type                           string        `json:"Type"`
			Name                           string        `json:"Name"`
			Description                    string        `json:"Description"`
			IfValidationFails              string        `json:"If Validation Fails"`
			IfValidatorNotApplicable       string        `json:"If Validator Not Applicable"`
			IssuancePhase                  string        `json:"Issuance Phase"`
			ApplyForAllCertificateProfiles bool          `json:"Apply for All Certificate Profiles"`
			AllowedDomainsList             []interface{} `json:"Allowed domains list"`
			AllowedDomainsListImportDate   string        `json:"Allowed domains list Import Date"`
			AllowedDomainsListSHA256       string        `json:"Allowed domains list SHA-256"`
		} `json:"AcmeAllowedDomains"`
		CaaValidator struct {
			ObjectType                  string        `json:"Object Type"`
			Version                     int           `json:"Version"`
			Type                        string        `json:"Type"`
			Name                        string        `json:"Name"`
			Description                 string        `json:"Description"`
			IfValidationFails           string        `json:"If Validation Fails"`
			IfValidatorNotApplicable    string        `json:"If Validator Not Applicable"`
			IssuancePhase               string        `json:"Issuance Phase"`
			ApplyForCertificateProfiles []interface{} `json:"Apply for Certificate Profiles"`
			Issuers                     []interface{} `json:"Issuers"`
			DNSResolver                 string        `json:"DNS Resolver"`
			LookupDNAMEs                bool          `json:"Lookup DNAMEs"`
			ValidateDNSSEC              bool          `json:"Validate DNSSEC"`
			DNSSECTrustAnchor           string        `json:"DNSSEC Trust Anchor"`
			FailOnLookupError           bool          `json:"Fail on Lookup Error"`
			TopLevelDomainsToIgnore     []interface{} `json:"Top-Level Domains to Ignore"`
			DomainsToIgnore             []interface{} `json:"Domains to Ignore"`
			IgnoredCriticalPropertyTags []interface{} `json:"Ignored Critical Property Tags"`
			DNSLookupRetryDelay         int           `json:"DNS Lookup Retry Delay"`
			MaximumLookupRecursions     int           `json:"Maximum Lookup Recursions"`
			IODEFSettings               struct {
				UseIODEFEMail bool `json:"Use IODEF E-Mail"`
				UseIODEFWeb   bool `json:"Use IODEF Web"`
			} `json:"IODEF Settings"`
		} `json:"CaaValidator"`
		PublicKeyBlock struct {
			ObjectType                  string        `json:"Object Type"`
			Version                     int           `json:"Version"`
			Type                        string        `json:"Type"`
			Name                        string        `json:"Name"`
			Description                 string        `json:"Description"`
			IfValidationFails           string        `json:"If Validation Fails"`
			IfValidatorNotApplicable    string        `json:"If Validator Not Applicable"`
			IssuancePhase               string        `json:"Issuance Phase"`
			ApplyForCertificateProfiles []interface{} `json:"Apply for Certificate Profiles"`
			KeyAlgorithms               []string      `json:"Key Algorithms"`
		} `json:"PublicKeyBlock"`
		RSAValidator struct {
			ObjectType                     string      `json:"Object Type"`
			Version                        int         `json:"Version"`
			Type                           string      `json:"Type"`
			Name                           string      `json:"Name"`
			Description                    string      `json:"Description"`
			IfValidationFails              string      `json:"If Validation Fails"`
			IfValidatorNotApplicable       string      `json:"If Validator Not Applicable"`
			IssuancePhase                  string      `json:"Issuance Phase"`
			ApplyForCertificateProfiles    []string    `json:"Apply for Certificate Profiles"`
			KeyValidatorSettingsTemplate   string      `json:"Key Validator Settings Template"`
			BitLengths                     []string    `json:"Bit Lengths"`
			OddPublicKeyExponentOnly       bool        `json:"Odd Public Key Exponent Only"`
			PublicKeyExponentMinimum       interface{} `json:"Public Key Exponent Minimum"`
			PublicKeyExponentMaximum       interface{} `json:"Public Key Exponent Maximum"`
			OddPublicKeyModulusOnly        bool        `json:"Odd Public Key Modulus Only"`
			DoNotAllowPowerOfPrimeModulus  bool        `json:"Do Not Allow Power of Prime Modulus"`
			DoNotAllowROCAWeakKeys         bool        `json:"Do Not Allow ROCA Weak Keys"`
			PublicKeyModulusSmallestFactor interface{} `json:"Public Key Modulus Smallest Factor"`
			PublicKeyModulusMinimum        interface{} `json:"Public Key Modulus Minimum"`
			PublicKeyModulusMaximum        interface{} `json:"Public Key Modulus Maximum"`
		} `json:"RSA Validator"`
		DomainBlocklistValidator struct {
			ObjectType                     string        `json:"Object Type"`
			Version                        int           `json:"Version"`
			Type                           string        `json:"Type"`
			Name                           string        `json:"Name"`
			Description                    string        `json:"Description"`
			IfValidationFails              string        `json:"If Validation Fails"`
			IfValidatorNotApplicable       string        `json:"If Validator Not Applicable"`
			IssuancePhase                  string        `json:"Issuance Phase"`
			ApplyForAllCertificateProfiles bool          `json:"Apply for All Certificate Profiles"`
			Blacklist                      []interface{} `json:"Blacklist"`
			BlacklistImportDate            string        `json:"Blacklist Import Date"`
			BlacklistSHA256                string        `json:"Blacklist SHA-256"`
			BlacklistChecks                []string      `json:"Blacklist Checks"`
			BlacklistNormalizations        []interface{} `json:"Blacklist Normalizations"`
		} `json:"Domain-Blocklist-Validator"`
		ZLintPreSignValidation struct {
			ObjectType                  string        `json:"Object Type"`
			Version                     int           `json:"Version"`
			Type                        string        `json:"Type"`
			Name                        string        `json:"Name"`
			Description                 string        `json:"Description"`
			IfValidationFails           string        `json:"If Validation Fails"`
			IfValidatorNotApplicable    string        `json:"If Validator Not Applicable"`
			IssuancePhase               string        `json:"Issuance Phase"`
			ApplyForCertificateProfiles []interface{} `json:"Apply for Certificate Profiles"`
			FailOnScriptErrorCode       bool          `json:"Fail on Script Error Code"`
			FailOnOutputToStandardError bool          `json:"Fail on Output to Standard Error"`
			LogStandardError            bool          `json:"Log Standard Error"`
			LogStandardOut              bool          `json:"Log Standard Out"`
			ExternalCommand             string        `json:"External Command"`
		} `json:"ZLint-Pre-SignValidation"`
		DomainValidatorTreasury struct {
			ObjectType                     string        `json:"Object Type"`
			Version                        int           `json:"Version"`
			Type                           string        `json:"Type"`
			Name                           string        `json:"Name"`
			Description                    string        `json:"Description"`
			IfValidationFails              string        `json:"If Validation Fails"`
			IfValidatorNotApplicable       string        `json:"If Validator Not Applicable"`
			IssuancePhase                  string        `json:"Issuance Phase"`
			ApplyForAllCertificateProfiles bool          `json:"Apply for All Certificate Profiles"`
			KeyValidatorSettingsTemplate   string        `json:"Key Validator Settings Template"`
			BitLengths                     []interface{} `json:"Bit Lengths"`
			OddPublicKeyExponentOnly       bool          `json:"Odd Public Key Exponent Only"`
			PublicKeyExponentMinimum       interface{}   `json:"Public Key Exponent Minimum"`
			PublicKeyExponentMaximum       interface{}   `json:"Public Key Exponent Maximum"`
			OddPublicKeyModulusOnly        bool          `json:"Odd Public Key Modulus Only"`
			DoNotAllowPowerOfPrimeModulus  bool          `json:"Do Not Allow Power of Prime Modulus"`
			DoNotAllowROCAWeakKeys         bool          `json:"Do Not Allow ROCA Weak Keys"`
			PublicKeyModulusSmallestFactor interface{}   `json:"Public Key Modulus Smallest Factor"`
			PublicKeyModulusMinimum        interface{}   `json:"Public Key Modulus Minimum"`
			PublicKeyModulusMaximum        interface{}   `json:"Public Key Modulus Maximum"`
		} `json:"Domain-Validator-Treasury"`
	} `json:"validators"`
	AdminPreferences struct {
		D4465831Da58A26Cb1F42Cb6Ff5B8A03Fb17 struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     int         `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"0499d4465831da58a26cb1f42cb6ff5b8a03fb17"`
		F826Cfd9Dc03B342D7Def7Ef8Da316Bd0Ee63740 struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"f826cfd9dc03b342d7def7ef8da316bd0ee63740"`
		B3B050Dfee2E3562A7836Cd1D8E829Bc44C1Bbf struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     interface{} `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"9b3b050dfee2e3562a7836cd1d8e829bc44c1bbf"`
		A51E87Feb6B1835F159F6682Dfad8C4400Fa908 struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"9a51e87feb6b1835f159f6682dfad8c4400fa908"`
		E771D110F56Cc0E070420408216229Ece1C5B005 struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     interface{} `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"e771d110f56cc0e070420408216229ece1c5b005"`
		Eb1Abfa12C8Dcc9A51E0Ca5Bbfe21A334740A77D struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"eb1abfa12c8dcc9a51e0ca5bbfe21a334740a77d"`
		E490De132D23A023409C760A229271D7266Fb struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     int         `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"004e490de132d23a023409c760a229271d7266fb"`
		C75B5E7A33A3C6824A7Fc4A4F6A7Ba7D906F2A struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"03c75b5e7a33a3c6824a7fc4a4f6a7ba7d906f2a"`
		Da7B533091Af7329A809A900D7180E69D178 struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"2094da7b533091af7329a809a900d7180e69d178"`
		Default struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     interface{} `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"default"`
		D4B7B1E793Aa139F5D0C8F5E03C01449Cbedb1D struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"0d4b7b1e793aa139f5d0c8f5e03c01449cbedb1d"`
		Db66193Ae50Fb87F3B019Dfd355D32224487Ccc8 struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"db66193ae50fb87f3b019dfd355d32224487ccc8"`
		F491Cb98B1A50Af21310Bdd8538A1F07Ca94A65B struct {
			ObjectType                      string      `json:"Object Type"`
			Version                         int         `json:"Version"`
			Name                            interface{} `json:"Name"`
			PreferredLanguage               string      `json:"Preferred Language"`
			SecondaryLanguage               string      `json:"Secondary Language"`
			EntriesPerPage                  int         `json:"Entries per Page"`
			LogEntriesPerPage               int         `json:"Log Entries per Page"`
			Theme                           string      `json:"Theme"`
			LastProfile                     string      `json:"Last Profile"`
			LastFilterMode                  string      `json:"Last Filter Mode"`
			LastLogFilterMode               string      `json:"Last Log Filter Mode"`
			FrontPageCAStatus               bool        `json:"Front Page CA Status"`
			FrontPagePublisherQueueStatus   bool        `json:"Front Page Publisher Queue Status"`
			PreferredRALanguage             interface{} `json:"Preferred RA Language"`
			PreferredRAStyleID              interface{} `json:"Preferred RA Style ID"`
			ConfigurationCheckerOnFrontPage bool        `json:"Configuration Checker on Front Page"`
		} `json:"f491cb98b1a50af21310bdd8538a1f07ca94a65b"`
	} `json:"admin-preferences"`
	TrustedOauthProviders struct {
		AzureC2 struct {
			ObjectType            string `json:"Object Type"`
			Version               int    `json:"Version"`
			Name                  string `json:"Name"`
			ClientSecret          string `json:"Client Secret"`
			Type                  int    `json:"Type"`
			Url                   string `json:"Url"`
			TokenUrl              string `json:"Token Url"`
			LogoutUrl             string `json:"Logout Url"`
			SkewLimitMs           int    `json:"Skew Limit (ms)"`
			Client                string `json:"Client"`
			Realm                 string `json:"Realm"`
			SCOPE                 string `json:"SCOPE"`
			AUDIENCE              string `json:"AUDIENCE"`
			AudienceCheckDisabled bool   `json:"Audience Check Disabled"`
			PublicKeyURL          string `json:"Public Key URL"`
			PublicKeys            []struct {
				KeyIdentifier string `json:"Key Identifier"`
				PublicKey     string `json:"Public Key"`
			} `json:"Public Keys"`
		} `json:"Azure-C2"`
	} `json:"trusted-oauth-providers"`
	ScepConfig struct {
		Device struct {
			ObjectType                            string      `json:"Object Type"`
			Version                               int         `json:"Version"`
			Name                                  string      `json:"Name"`
			OperationalMode                       string      `json:"Operational Mode"`
			ReturnCAChainInGetCACertResponse      bool        `json:"Return CA Chain In GetCACert response"`
			AllowLegacyDigestAlgorithmsInResponse bool        `json:"Allow Legacy Digest Algorithms in Response"`
			RAEndEntityProfile                    string      `json:"RA End Entity Profile"`
			RACertificateProfile                  string      `json:"RA Certificate Profile"`
			RACAName                              string      `json:"RA CA Name"`
			RAAuthenticationPassword              string      `json:"RA Authentication Password"`
			RANameGenerationParameters            interface{} `json:"RA Name Generation Parameters"`
		} `json:"device"`
		Firewall struct {
			ObjectType                            string      `json:"Object Type"`
			Version                               int         `json:"Version"`
			Name                                  string      `json:"Name"`
			OperationalMode                       string      `json:"Operational Mode"`
			ReturnCAChainInGetCACertResponse      bool        `json:"Return CA Chain In GetCACert response"`
			AllowLegacyDigestAlgorithmsInResponse bool        `json:"Allow Legacy Digest Algorithms in Response"`
			RAEndEntityProfile                    string      `json:"RA End Entity Profile"`
			RACertificateProfile                  string      `json:"RA Certificate Profile"`
			RACAName                              string      `json:"RA CA Name"`
			RAAuthenticationPassword              string      `json:"RA Authentication Password"`
			RANameGenerationParameters            interface{} `json:"RA Name Generation Parameters"`
		} `json:"firewall"`
		Intune struct {
			ObjectType                            string `json:"Object Type"`
			Version                               int    `json:"Version"`
			Name                                  string `json:"Name"`
			OperationalMode                       string `json:"Operational Mode"`
			ReturnCAChainInGetCACertResponse      bool   `json:"Return CA Chain In GetCACert response"`
			AllowLegacyDigestAlgorithmsInResponse bool   `json:"Allow Legacy Digest Algorithms in Response"`
			AuthenticateThroughMicrosoftIntune    string `json:"Authenticate through Microsoft Intune"`
			RAEndEntityProfile                    string `json:"RA End Entity Profile"`
			RACertificateProfile                  string `json:"RA Certificate Profile"`
			RACAName                              string `json:"RA CA Name"`
			RAAuthenticationPassword              string `json:"RA Authentication Password"`
			AzureApplicationID                    string `json:"Azure Application ID"`
			AzureApplicationAPISecret             string `json:"Azure Application API Secret"`
			IntuneTenant                          string `json:"Intune Tenant"`
		} `json:"intune"`
		Intunera struct {
			ObjectType                            string `json:"Object Type"`
			Version                               int    `json:"Version"`
			Name                                  string `json:"Name"`
			OperationalMode                       string `json:"Operational Mode"`
			ReturnCAChainInGetCACertResponse      bool   `json:"Return CA Chain In GetCACert response"`
			AllowLegacyDigestAlgorithmsInResponse bool   `json:"Allow Legacy Digest Algorithms in Response"`
			AuthenticateThroughMicrosoftIntune    string `json:"Authenticate through Microsoft Intune"`
			RAEndEntityProfile                    string `json:"RA End Entity Profile"`
			RACertificateProfile                  string `json:"RA Certificate Profile"`
			RACAName                              string `json:"RA CA Name"`
			RAAuthenticationPassword              string `json:"RA Authentication Password"`
			AzureApplicationID                    string `json:"Azure Application ID"`
			AzureApplicationAPISecret             string `json:"Azure Application API Secret"`
			IntuneTenant                          string `json:"Intune Tenant"`
		} `json:"intunera"`
		Primekeyqa struct {
			ObjectType                            string `json:"Object Type"`
			Version                               int    `json:"Version"`
			Name                                  string `json:"Name"`
			OperationalMode                       string `json:"Operational Mode"`
			ReturnCAChainInGetCACertResponse      bool   `json:"Return CA Chain In GetCACert response"`
			AllowLegacyDigestAlgorithmsInResponse bool   `json:"Allow Legacy Digest Algorithms in Response"`
			AuthenticateThroughMicrosoftIntune    string `json:"Authenticate through Microsoft Intune"`
			RAEndEntityProfile                    string `json:"RA End Entity Profile"`
			RACertificateProfile                  string `json:"RA Certificate Profile"`
			RACAName                              string `json:"RA CA Name"`
			RAAuthenticationPassword              string `json:"RA Authentication Password"`
			AzureApplicationID                    string `json:"Azure Application ID"`
			AzureApplicationAPISecret             string `json:"Azure Application API Secret"`
			IntuneTenant                          string `json:"Intune Tenant"`
		} `json:"primekeyqa"`
		Router struct {
			ObjectType string `json:"Object Type"`
			Version    int    `json:"Version"`
			Name       string `json:"Name"`
		} `json:"router"`
		Ryan struct {
			ObjectType                            string `json:"Object Type"`
			Version                               int    `json:"Version"`
			Name                                  string `json:"Name"`
			OperationalMode                       string `json:"Operational Mode"`
			IncludeCACertificateInResponse        bool   `json:"Include CA Certificate in Response"`
			AllowLegacyDigestAlgorithmsInResponse bool   `json:"Allow Legacy Digest Algorithms in Response"`
			RAEndEntityProfile                    string `json:"RA End Entity Profile"`
			RACertificateProfile                  string `json:"RA Certificate Profile"`
			RACAName                              string `json:"RA CA Name"`
			RAAuthenticationPassword              string `json:"RA Authentication Password"`
		} `json:"ryan"`
		SCEPAlias struct {
			ObjectType string `json:"Object Type"`
			Version    int    `json:"Version"`
			Name       string `json:"Name"`
		} `json:"SCEPAlias"`
	} `json:"scep-config"`
	AvailableProtocols struct {
		AvailableProtocolConfiguration struct {
			ObjectType                              string `json:"Object Type"`
			Version                                 int    `json:"Version"`
			Name                                    string `json:"Name"`
			ACME                                    bool   `json:"ACME"`
			Certstore                               bool   `json:"Certstore"`
			CMP                                     bool   `json:"CMP"`
			CRLstore                                bool   `json:"CRLstore"`
			EST                                     bool   `json:"EST"`
			MSAE                                    bool   `json:"MSAE"`
			OCSP                                    bool   `json:"OCSP"`
			PublicWeb                               bool   `json:"Public Web"`
			SCEP                                    bool   `json:"SCEP"`
			RAWeb                                   bool   `json:"RA Web"`
			RESTCAManagement                        bool   `json:"REST CA Management"`
			RESTCertificateManagement               bool   `json:"REST Certificate Management"`
			RESTCryptoTokenManagement               bool   `json:"REST Crypto Token Management"`
			RESTEndEntityManagement                 bool   `json:"REST End Entity Management"`
			RESTEndEntityManagementV2               bool   `json:"REST End Entity Management V2"`
			RESTConfigdump                          bool   `json:"REST Configdump"`
			RESTCertificateManagementV2             bool   `json:"REST Certificate Management V2"`
			RESTSSHV1                               bool   `json:"REST SSH V1"`
			Webdist                                 bool   `json:"Webdist"`
			WebService                              bool   `json:"Web Service"`
			ITSCertificateManagement                bool   `json:"ITS Certificate Management"`
			CustomHeaderNameForRESTCallsFromBrowser string `json:"Custom header name for REST calls from browser"`
		} `json:"available-protocol-configuration"`
	} `json:"available-protocols"`
	Services struct {
		DailyCRLUpdate struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				CAsToCheck []string `json:"CAs to Check"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"DailyCRLUpdate"`
		AzureIntuneRevocation struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				MaxDownloads           string `json:"maxDownloads"`
				ProxyUser              string `json:"proxyUser"`
				IssuerDns              string `json:"issuerDns"`
				GraphApiVersion        string `json:"graphApiVersion"`
				IntuneAppId            string `json:"intuneAppId"`
				ProviderNameAndVersion string `json:"providerNameAndVersion"`
				ActiveDirectoryAuthUrl string `json:"activeDirectoryAuthUrl"`
				IntuneResourceUrl      string `json:"intuneResourceUrl"`
				ProxyPass              string `json:"proxyPass"`
				PeerAndBinding         string `json:"peerAndBinding"`
				AppKey                 string `json:"appKey"`
				ProxyHost              string `json:"proxyHost"`
				Tenant                 string `json:"tenant"`
				GraphResourceUrl       string `json:"graphResourceUrl"`
				AppId                  string `json:"appId"`
				ProxyPort              string `json:"proxyPort"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string   `json:"Description"`
			Active             bool     `json:"Active"`
			Hidden             bool     `json:"Hidden"`
			PinToSpecificNodes []string `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool     `json:"Run on All Nodes"`
		} `json:"azureIntuneRevocation"`
		PublishQueue struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				TimeBeforeExpiring int      `json:"Time Before Expiring"`
				TimeUnit           string   `json:"Time Unit"`
				SendToEndUsers     bool     `json:"Send to End Users"`
				UserEMailSubject   string   `json:"User E-mail Subject"`
				UserEMailMessage   string   `json:"User E-mail Message"`
				SendToAdmins       bool     `json:"Send to Admins"`
				AdminEMailSubject  string   `json:"Admin E-mail Subject"`
				AdminEMailMessage  string   `json:"Admin E-mail Message"`
				PublishersToCheck  []string `json:"Publishers to Check"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"PublishQueue"`
		As struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"as"`
		CrlDownload struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				CAsToCheck       []int `json:"CAs to Check"`
				MaxDownloadSize  int   `json:"Max Download Size"`
				IgnoreNextUpdate bool  `json:"Ignore nextUpdate"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"crlDownload"`
		PasswordExpire struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				TimeBeforeExpiring int    `json:"Time Before Expiring"`
				TimeUnit           string `json:"Time Unit"`
				SendToEndUsers     bool   `json:"Send to End Users"`
				UserEMailSubject   string `json:"User E-mail Subject"`
				UserEMailMessage   string `json:"User E-mail Message"`
				SendToAdmins       bool   `json:"Send to Admins"`
				AdminEMailSubject  string `json:"Admin E-mail Subject"`
				AdminEMailMessage  string `json:"Admin E-mail Message"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"passwordExpire"`
		HSMKeepalive struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"HSM-Keepalive"`
		OcspKeyUpdate struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				CAsToCheck                   []string `json:"CAs to Check"`
				PeerSystem                   string   `json:"Peer System"`
				TimeBeforeCertificateExpires int      `json:"Time Before Certificate Expires"`
				TimeUnitOfCertificateExpiry  string   `json:"Time Unit of Certificate Expiry"`
				RenewKeys                    bool     `json:"Renew Keys"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"ocspKeyUpdate"`
		OcspPreSign struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				CAsToCheck                          []string `json:"CAs to Check"`
				TimeBeforeExpiring                  int      `json:"Time Before Expiring"`
				TimeUnit                            string   `json:"Time Unit"`
				CertIdHashAlgorithm                 string   `json:"CertId Hash Algorithm"`
				GenerateResponsesForAllCertificates bool     `json:"Generate Responses for All Certificates"`
				UpdateExpiredResponsesOnly          bool     `json:"Update Expired Responses Only"`
				IssueFinalOCSPResponseEIDAS         bool     `json:"Issue Final OCSP Response (eIDAS)"`
				TimeBeforeCAExpires                 int      `json:"Time Before CA Expires"`
				TimeUnitBeforeCAExpires             string   `json:"Time Unit Before CA Expires"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"ocspPreSign"`
		PublishQueueAws struct {
			ObjectType       string `json:"Object Type"`
			Version          int    `json:"Version"`
			Name             string `json:"Name"`
			WorkerType       string `json:"Worker Type"`
			WorkerProperties struct {
				TimeBeforeExpiring int    `json:"Time Before Expiring"`
				TimeUnit           string `json:"Time Unit"`
				SendToEndUsers     bool   `json:"Send to End Users"`
				UserEMailSubject   string `json:"User E-mail Subject"`
				UserEMailMessage   string `json:"User E-mail Message"`
				SendToAdmins       bool   `json:"Send to Admins"`
				AdminEMailSubject  string `json:"Admin E-mail Subject"`
				AdminEMailMessage  string `json:"Admin E-mail Message"`
				PublishersToCheck  []int  `json:"Publishers to Check"`
			} `json:"Worker Properties"`
			IntervalType       string `json:"Interval Type"`
			IntervalProperties struct {
				PeriodicalValue int    `json:"Periodical Value"`
				PeriodicalUnit  string `json:"Periodical Unit"`
			} `json:"Interval Properties"`
			ActionType       string `json:"Action Type"`
			ActionProperties struct {
			} `json:"Action Properties"`
			Description        string        `json:"Description"`
			Active             bool          `json:"Active"`
			Hidden             bool          `json:"Hidden"`
			PinToSpecificNodes []interface{} `json:"Pin to Specific Nodes"`
			RunOnAllNodes      bool          `json:"Run on All Nodes"`
		} `json:"PublishQueue-aws"`
	} `json:"services"`
	CmpConfig struct {
		Gpp struct {
			ObjectType                                 string `json:"Object Type"`
			Version                                    int    `json:"Version"`
			Name                                       string `json:"Name"`
			CMPOperationalModeIsRA                     bool   `json:"CMP Operational Mode is RA"`
			CMPAuthenticationModule                    string `json:"CMP Authentication Module"`
			CMPAuthenticationModuleParameters          string `json:"CMP Authentication Module Parameters"`
			ExtractUsernameComponent                   string `json:"Extract Username Component"`
			RANameGenerationPrefix                     string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix                    string `json:"RA Name Generation Postfix"`
			VendorCertificateMode                      bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                            string `json:"List of Vendor CAs"`
			CMPResponseProtection                      string `json:"CMP Response Protection"`
			CMPResponseAdditionalCACertificates        string `json:"CMP Response Additional CA Certificates"`
			CMPResponseIssuingCACertificateAtIndex0    bool   `json:"CMP Response Issuing CA Certificate at index 0"`
			PKIMessageResponseAdditionalCACertificates string `json:"PKI Message Response Additional CA Certificates"`
			DefaultCA                                  string `json:"Default CA"`
			AutomaticKeyUpdate                         bool   `json:"Automatic Key Update"`
			CertificateRenewalWithSameKeys             bool   `json:"Certificate Renewal with Same Keys"`
			AllowServerGeneratedKeys                   bool   `json:"Allow Server Generated Keys"`
			TrustedCertificatesPath                    string `json:"Trusted Certificates Path"`
			ExtendedValidation                         bool   `json:"Extended Validation"`
		} `json:"3gpp"`
		Gnodeb struct {
			ObjectType                                 string `json:"Object Type"`
			Version                                    int    `json:"Version"`
			Name                                       string `json:"Name"`
			CMPOperationalModeIsRA                     bool   `json:"CMP Operational Mode is RA"`
			CMPAuthenticationModule                    string `json:"CMP Authentication Module"`
			CMPAuthenticationModuleParameters          string `json:"CMP Authentication Module Parameters"`
			ExtractUsernameComponent                   string `json:"Extract Username Component"`
			RANameGenerationPrefix                     string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix                    string `json:"RA Name Generation Postfix"`
			VendorCertificateMode                      bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                            string `json:"List of Vendor CAs"`
			CMPResponseProtection                      string `json:"CMP Response Protection"`
			CMPResponseAdditionalCACertificates        string `json:"CMP Response Additional CA Certificates"`
			CMPResponseIssuingCACertificateAtIndex0    bool   `json:"CMP Response Issuing CA Certificate at index 0"`
			PKIMessageResponseAdditionalCACertificates string `json:"PKI Message Response Additional CA Certificates"`
			DefaultCA                                  string `json:"Default CA"`
			AutomaticKeyUpdate                         bool   `json:"Automatic Key Update"`
			CertificateRenewalWithSameKeys             bool   `json:"Certificate Renewal with Same Keys"`
			AllowServerGeneratedKeys                   bool   `json:"Allow Server Generated Keys"`
			TrustedCertificatesPath                    string `json:"Trusted Certificates Path"`
			ExtendedValidation                         bool   `json:"Extended Validation"`
		} `json:"gnodeb"`
		Gnodebra struct {
			ObjectType                                              string `json:"Object Type"`
			Version                                                 int    `json:"Version"`
			Name                                                    string `json:"Name"`
			CMPOperationalModeIsRA                                  bool   `json:"CMP Operational Mode is RA"`
			CMPAuthenticationModule                                 string `json:"CMP Authentication Module"`
			CMPAuthenticationModuleParameters                       string `json:"CMP Authentication Module Parameters"`
			RAVerifyProofOfPossession                               bool   `json:"RA Verify Proof of Possession"`
			RANameGenerationScheme                                  string `json:"RA Name Generation Scheme"`
			RANameGenerationSchemeParameters                        string `json:"RA Name Generation Scheme Parameters"`
			RANameGenerationPrefix                                  string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix                                 string `json:"RA Name Generation Postfix"`
			PasswordGenerationParameter                             string `json:"Password Generation Parameter"`
			CustomCertificateSerialNumber                           bool   `json:"Custom Certificate Serial Number"`
			RAEndEntityProfile                                      string `json:"RA End Entity Profile"`
			RACertificateProfile                                    string `json:"RA Certificate Profile"`
			RACAName                                                string `json:"RA CA Name"`
			VerificationsInEndEntityCertificateAuthenticationModule bool   `json:"Verifications in EndEntityCertificate Authentication Module"`
			CMPResponseProtection                                   string `json:"CMP Response Protection"`
			CMPResponseAdditionalCACertificates                     string `json:"CMP Response Additional CA Certificates"`
			CMPResponseIssuingCACertificateAtIndex0                 bool   `json:"CMP Response Issuing CA Certificate at index 0"`
			PKIMessageResponseAdditionalCACertificates              string `json:"PKI Message Response Additional CA Certificates"`
			DefaultCA                                               string `json:"Default CA"`
			AutomaticKeyUpdate                                      bool   `json:"Automatic Key Update"`
			CertificateRenewalWithSameKeys                          bool   `json:"Certificate Renewal with Same Keys"`
			AllowServerGeneratedKeys                                bool   `json:"Allow Server Generated Keys"`
			TrustedCertificatesPath                                 string `json:"Trusted Certificates Path"`
			ExtendedValidation                                      bool   `json:"Extended Validation"`
		} `json:"gnodebra"`
		Device struct {
			ObjectType                                              string `json:"Object Type"`
			Version                                                 int    `json:"Version"`
			Name                                                    string `json:"Name"`
			CMPOperationalModeIsRA                                  bool   `json:"CMP Operational Mode is RA"`
			CMPAuthenticationModule                                 string `json:"CMP Authentication Module"`
			CMPAuthenticationModuleParameters                       string `json:"CMP Authentication Module Parameters"`
			RAVerifyProofOfPossession                               bool   `json:"RA Verify Proof of Possession"`
			RANameGenerationScheme                                  string `json:"RA Name Generation Scheme"`
			RANameGenerationSchemeParameters                        string `json:"RA Name Generation Scheme Parameters"`
			RANameGenerationPrefix                                  string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix                                 string `json:"RA Name Generation Postfix"`
			PasswordGenerationParameter                             string `json:"Password Generation Parameter"`
			CustomCertificateSerialNumber                           bool   `json:"Custom Certificate Serial Number"`
			RAEndEntityProfile                                      string `json:"RA End Entity Profile"`
			RACertificateProfile                                    string `json:"RA Certificate Profile"`
			RACAName                                                string `json:"RA CA Name"`
			VerificationsInEndEntityCertificateAuthenticationModule bool   `json:"Verifications in EndEntityCertificate Authentication Module"`
			CMPResponseProtection                                   string `json:"CMP Response Protection"`
			CMPResponseAdditionalCACertificates                     string `json:"CMP Response Additional CA Certificates"`
			CMPResponseIssuingCACertificateAtIndex0                 bool   `json:"CMP Response Issuing CA Certificate at index 0"`
			PKIMessageResponseAdditionalCACertificates              string `json:"PKI Message Response Additional CA Certificates"`
			DefaultCA                                               string `json:"Default CA"`
			AutomaticKeyUpdate                                      bool   `json:"Automatic Key Update"`
			CertificateRenewalWithSameKeys                          bool   `json:"Certificate Renewal with Same Keys"`
			AllowServerGeneratedKeys                                bool   `json:"Allow Server Generated Keys"`
			TrustedCertificatesPath                                 string `json:"Trusted Certificates Path"`
			ExtendedValidation                                      bool   `json:"Extended Validation"`
		} `json:"device"`
		Carl struct {
			ObjectType                                 string `json:"Object Type"`
			Version                                    int    `json:"Version"`
			Name                                       string `json:"Name"`
			CMPOperationalModeIsRA                     bool   `json:"CMP Operational Mode is RA"`
			CMPAuthenticationModule                    string `json:"CMP Authentication Module"`
			CMPAuthenticationModuleParameters          string `json:"CMP Authentication Module Parameters"`
			ExtractUsernameComponent                   string `json:"Extract Username Component"`
			RANameGenerationPrefix                     string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix                    string `json:"RA Name Generation Postfix"`
			VendorCertificateMode                      bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                            string `json:"List of Vendor CAs"`
			CMPResponseProtection                      string `json:"CMP Response Protection"`
			CMPResponseAdditionalCACertificates        string `json:"CMP Response Additional CA Certificates"`
			CMPResponseIssuingCACertificateAtIndex0    bool   `json:"CMP Response Issuing CA Certificate at index 0"`
			PKIMessageResponseAdditionalCACertificates string `json:"PKI Message Response Additional CA Certificates"`
			DefaultCA                                  string `json:"Default CA"`
			AutomaticKeyUpdate                         bool   `json:"Automatic Key Update"`
			CertificateRenewalWithSameKeys             bool   `json:"Certificate Renewal with Same Keys"`
			AllowServerGeneratedKeys                   bool   `json:"Allow Server Generated Keys"`
			TrustedCertificatesPath                    string `json:"Trusted Certificates Path"`
			ExtendedValidation                         bool   `json:"Extended Validation"`
		} `json:"carl"`
	} `json:"cmp-config"`
	CertificateProfiles struct {
		DigitalSignature20483Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"DigitalSignature-2048-3y"`
		SubCAITECG1 struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed         bool          `json:"Path Length Constraint Used"`
			PathLengthConstraint             int           `json:"Path Length Constraint"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"SubCA-IT-EC-G1"`
		TlsServerAuthApproval2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CertificateExtensions            []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthApproval2d"`
		TlsServerAuth struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuth"`
		Workstation2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CertificateExtensions            []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"workstation2d"`
		Signer3Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			PrivateKeyUsagePeriodLengthUsed  bool          `json:"Private Key Usage Period Length Used"`
			PrivateKeyUsagePeriodLength      string        `json:"Private Key Usage Period Length"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"signer-3y"`
		SubCAMsIoTG1 struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed         bool          `json:"Path Length Constraint Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"SubCA-MsIoT-G1"`
		SignerCodeSigning3Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"signerCodeSigning-3y"`
		AlteryxRootCa struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed      bool          `json:"Authority Key Identifier Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem []string      `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"AlteryxRootCa"`
		TlsPaloAltoClientAuth struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsPaloAltoClientAuth"`
		IntuneClientAuth1Y struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed                    bool          `json:"Basic Constraints Used"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed                    bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                        []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable        bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed               bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed                bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings                 bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed          bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined        bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined                bool          `json:"OCSP Issuer URL Is CA Defined"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CVCAccessRightsInspectionSystem         []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"intuneClientAuth-1y"`
		Workstation10M struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"workstation10m"`
		TlsServerMgmtCA struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			CRLDistributionPointURI          string        `json:"CRL Distribution Point URI"`
			CRLIssuer                        string        `json:"CRL Issuer"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPServiceLocatorURL            string        `json:"OCSP Service Locator URL"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			OCSPIssuerURLs                   []string      `json:"OCSP Issuer URLs"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerMgmtCA"`
		HashiCorpSubCAG1 struct {
			ObjectType                                   string        `json:"Object Type"`
			Version                                      int           `json:"Version"`
			Name                                         string        `json:"Name"`
			Type                                         string        `json:"Type"`
			AvailableKeyAlgorithms                       []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                      []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                          []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                           string        `json:"Signature Algorithm"`
			Validity                                     string        `json:"Validity"`
			Description                                  string        `json:"Description"`
			AllowSubjectDNOverrideByEndEntityInformation bool          `json:"Allow Subject DN Override by End Entity Information"`
			OverridableExtensionOIDs                     []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs                  []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride                        bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed                     bool          `json:"Path Length Constraint Used"`
			KeyUsage                                     []string      `json:"Key Usage"`
			ExtendedKeyUsage                             []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed                   bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed                    bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed                     bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings                      bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed               bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined             bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined                     bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem              []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                                  bool          `json:"LDAP DN Order"`
			AvailableCAs                                 []string      `json:"Available CAs"`
			AccountBindingNamespace                      []interface{} `json:"Account Binding Namespace"`
		} `json:"HashiCorp-SubCA-G1"`
		WorkstationEFS3Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CertificateExtensions            []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"workstationEFS-3y"`
		IotDevice struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"iot-device"`
		RootCAITECG1 struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed      bool          `json:"Authority Key Identifier Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"RootCA-IT-EC-G1"`
		DomainController struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CertificateExtensions            []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"domainController"`
		TlsServerAuthAcme2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthAcme2d"`
		IkeIntermediate struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed                    bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                        []string      `json:"Extended Key Usage"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CVCAccessRightsInspectionSystem         []string      `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"ike-intermediate"`
		TlsServerClientAuth struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerClientAuth"`
		TlsServerAuthAcme90D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthAcme90d"`
		SolitudeDomainController struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"solitude-domainController"`
		AdminMgmtCA struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"adminMgmtCA"`
		Xcert1 struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CertificateExtensions            []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"xcert1"`
		TlsPaloAltoClientAuth90D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsPaloAltoClientAuth-90d"`
		SubCADPodG1 struct {
			ObjectType                  string        `json:"Object Type"`
			Version                     int           `json:"Version"`
			Name                        string        `json:"Name"`
			Type                        string        `json:"Type"`
			AvailableKeyAlgorithms      []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves     []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths         []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm          string        `json:"Signature Algorithm"`
			Validity                    string        `json:"Validity"`
			Description                 string        `json:"Description"`
			OverridableExtensionOIDs    []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride       bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed    bool          `json:"Path Length Constraint Used"`
			KeyUsage                    []string      `json:"Key Usage"`
			ExtendedKeyUsage            []interface{} `json:"Extended Key Usage"`
			CertificatePoliciesUsed     bool          `json:"Certificate Policies Used"`
			CertificatePolicies         []struct {
				PolicyID    string `json:"Policy ID"`
				QualifierID string `json:"Qualifier ID"`
				Qualifier   string `json:"Qualifier"`
			} `json:"Certificate Policies"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"SubCA-DPod-G1"`
		Encryption20483Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"Encryption-2048-3y"`
		TlsClientAuthApproval struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsClientAuthApproval"`
		RootCAITG1 struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed      bool          `json:"Authority Key Identifier Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"RootCA-IT-G1"`
		TlsClientAuth struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []string      `json:"Account Binding Namespace"`
		} `json:"tlsClientAuth"`
		SshHostRsa1Y struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            int           `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed            bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                []string      `json:"Extended Key Usage"`
			CVCAccessRightsInspectionSystem []string      `json:"CVC Access Rights (Inspection System)"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"sshHostRsa1y"`
		SolnaWorkstation2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []string      `json:"Account Binding Namespace"`
		} `json:"solna-workstation2d"`
		IstioAuth3D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			UseCertificateStorage            bool          `json:"Use Certificate Storage"`
			StoreCertificateData             bool          `json:"Store Certificate Data"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			CertificatePoliciesUsed          bool          `json:"Certificate Policies Used"`
			CertificatePolicies              []interface{} `json:"Certificate Policies"`
			SubjectAlternativeNameCritical   bool          `json:"Subject Alternative Name Critical"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"istioAuth-3d"`
		New struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed            bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                []string      `json:"Extended Key Usage"`
			CVCAccessRightsInspectionSystem []string      `json:"CVC Access Rights (Inspection System)"`
			AvailableCAs                    []interface{} `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"new"`
		CSCA30Y struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride                   bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed              bool          `json:"Authority Key Identifier Used"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsage                        []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed              bool          `json:"Subject Alternative Name Used"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CVCAccessRightsInspectionSystem         []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"CSCA-30y"`
		CarlTlsClientAuth7D struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed                    bool          `json:"Basic Constraints Used"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed                    bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                        []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable        bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed               bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed                bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings                 bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed          bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined        bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined                bool          `json:"OCSP Issuer URL Is CA Defined"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CVCAccessRightsInspectionSystem         []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			SubjectDNSubsetUsed                     bool          `json:"Subject DN Subset Used"`
			SubjectDNSubset                         []string      `json:"Subject DN Subset"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"carl-tlsClientAuth-7d"`
		OcspSigner struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed            bool          `json:"Basic Constraints Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed            bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			OCSPNoCheck                     bool          `json:"OCSP No Check"`
			CVCAccessRightsInspectionSystem []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"ocspSigner"`
		RootCALDAP struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed      bool          `json:"Authority Key Identifier Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem []interface{} `json:"CVC Access Rights (Inspection System)"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"RootCA-LDAP"`
		SubCAMsIoTECG1 struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed         bool          `json:"Path Length Constraint Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"SubCA-MsIoT-EC-G1"`
		SshCaRsa5Y struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			CVCAccessRightsInspectionSystem []string      `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"sshCaRsa-5y"`
		CodeSigning1Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"codeSigning-1y"`
		AllowExtensionOverride struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			AllowExtensionOverride           bool          `json:"Allow Extension Override"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"allowExtensionOverride"`
		Authentication20483Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"Authentication-2048-3y"`
		Authentication2048Approve1Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"Authentication-2048-Approve-1y"`
		TlsServerAuthAcme30D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthAcme30d"`
		TlsServerAuthEC struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed                    bool          `json:"Basic Constraints Used"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed                    bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                        []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable        bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed               bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed                bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings                 bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed          bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined        bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined                bool          `json:"OCSP Issuer URL Is CA Defined"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CVCAccessRightsInspectionSystem         []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuth-EC"`
		SolnaTlsServerAuthAcme2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []string      `json:"Account Binding Namespace"`
		} `json:"solna-tlsServerAuthAcme2d"`
		VpnFirewall struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"vpnFirewall"`
		RemoteDesktopAuth struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"remoteDesktopAuth"`
		SubCA05ITG1 struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed         bool          `json:"Path Length Constraint Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"SubCA05-IT-G1"`
		SubCAITG1 struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed         bool          `json:"Path Length Constraint Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"SubCA-IT-G1"`
		TlsServerClientAuthApproval struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerClientAuthApproval"`
		SolitudeAuthenticationMsae struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed                    bool          `json:"Basic Constraints Used"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed                    bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                        []string      `json:"Extended Key Usage"`
			IssuerAlternativeNameUsed               bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed                bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings                 bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed          bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined        bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined                bool          `json:"OCSP Issuer URL Is CA Defined"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CertificateExtensions                   []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem         []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"solitude-authentication-msae"`
		TlsServerAuthApproval struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthApproval"`
		CarlTlsClientAuth1Y struct {
			ObjectType                              string        `json:"Object Type"`
			Version                                 int           `json:"Version"`
			Name                                    string        `json:"Name"`
			Type                                    string        `json:"Type"`
			AvailableKeyAlgorithms                  []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves                 []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths                     []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm                      string        `json:"Signature Algorithm"`
			Validity                                string        `json:"Validity"`
			Description                             string        `json:"Description"`
			OverridableExtensionOIDs                []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs             []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed                    bool          `json:"Basic Constraints Used"`
			KeyUsage                                []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed                    bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                        []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable        bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed               bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed                bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings                 bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed          bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined        bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined                bool          `json:"OCSP Issuer URL Is CA Defined"`
			MicrosoftObjectSidSecurityExtensionUsed bool          `json:"Microsoft ObjectSid Security Extension Used"`
			CVCAccessRightsInspectionSystem         []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                             bool          `json:"LDAP DN Order"`
			SubjectDNSubsetUsed                     bool          `json:"Subject DN Subset Used"`
			SubjectDNSubset                         []string      `json:"Subject DN Subset"`
			AvailableCAs                            []string      `json:"Available CAs"`
			AccountBindingNamespace                 []interface{} `json:"Account Binding Namespace"`
		} `json:"carl-tlsClientAuth-1y"`
		SolitudeUserEcdsaMsae struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"solitude-user-ecdsa-msae"`
		RgRemoteSvcServerClientAuth1Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"rgRemoteSvcServerClientAuth-1y"`
		TlsClientAuth2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsClientAuth-2d"`
		ClevelandWorkstation2D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []string      `json:"Account Binding Namespace"`
		} `json:"cleveland-workstation2d"`
		VpnFirewall10M struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"vpnFirewall10m"`
		SignerPdf3Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"signerPdf-3y"`
		AdminNoCDPMgmtCA struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"AdminNoCDPMgmtCA"`
		IotCa struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []string      `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"iot-ca"`
		RemoteDesktopAuth3Mo struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"remoteDesktopAuth-3mo"`
		RootCAMsIoTG1 struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed      bool          `json:"Authority Key Identifier Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"RootCA-MsIoT-G1"`
		TlsServerAuthAcme8D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthAcme8d"`
		TlsServerAuthAcme60D struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsServerAuthAcme60d"`
		Workstation struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"workstation"`
		SignerTimeStampAuthority8Y struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsageCritical         bool          `json:"Extended Key Usage Critical"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"signerTimeStampAuthority-8y"`
		ManagementCA struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride            bool          `json:"Allow Validity Override"`
			PathLengthConstraintUsed         bool          `json:"Path Length Constraint Used"`
			AuthorityKeyIdentifierUsed       bool          `json:"Authority Key Identifier Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsage                 []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"managementCA"`
		TlsPeerConnector struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed       bool          `json:"Subject Alternative Name Used"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"tlsPeerConnector"`
		RootCAMsIoTECG1 struct {
			ObjectType                      string        `json:"Object Type"`
			Version                         int           `json:"Version"`
			Name                            string        `json:"Name"`
			Type                            string        `json:"Type"`
			AvailableKeyAlgorithms          []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves         []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths             []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm              string        `json:"Signature Algorithm"`
			Validity                        string        `json:"Validity"`
			Description                     string        `json:"Description"`
			OverridableExtensionOIDs        []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs     []interface{} `json:"Non-overridable Extension OIDs"`
			AllowValidityOverride           bool          `json:"Allow Validity Override"`
			AuthorityKeyIdentifierUsed      bool          `json:"Authority Key Identifier Used"`
			KeyUsage                        []string      `json:"Key Usage"`
			ExtendedKeyUsage                []interface{} `json:"Extended Key Usage"`
			SubjectAlternativeNameUsed      bool          `json:"Subject Alternative Name Used"`
			IssuerAlternativeNameUsed       bool          `json:"Issuer Alternative Name Used"`
			CVCAccessRightsInspectionSystem []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                     bool          `json:"LDAP DN Order"`
			AvailableCAs                    []string      `json:"Available CAs"`
			AccountBindingNamespace         []interface{} `json:"Account Binding Namespace"`
		} `json:"RootCA-MsIoT-EC-G1"`
		SolitudeWorkstationMsae struct {
			ObjectType                       string        `json:"Object Type"`
			Version                          int           `json:"Version"`
			Name                             string        `json:"Name"`
			Type                             string        `json:"Type"`
			AvailableKeyAlgorithms           []string      `json:"Available Key Algorithms"`
			AvailableEllipticCurves          []string      `json:"Available Elliptic Curves"`
			AvailableBitLengths              []int         `json:"Available Bit Lengths"`
			SignatureAlgorithm               string        `json:"Signature Algorithm"`
			Validity                         string        `json:"Validity"`
			Description                      string        `json:"Description"`
			OverridableExtensionOIDs         []interface{} `json:"Overridable Extension OIDs"`
			NonOverridableExtensionOIDs      []interface{} `json:"Non-overridable Extension OIDs"`
			BasicConstraintsUsed             bool          `json:"Basic Constraints Used"`
			KeyUsage                         []string      `json:"Key Usage"`
			ExtendedKeyUsageUsed             bool          `json:"Extended Key Usage Used"`
			ExtendedKeyUsage                 []string      `json:"Extended Key Usage"`
			SubjectAlternativeNameSearchable bool          `json:"Subject Alternative Name Searchable"`
			IssuerAlternativeNameUsed        bool          `json:"Issuer Alternative Name Used"`
			CRLDistributionPointUsed         bool          `json:"CRL Distribution Point Used"`
			CRLUseCADefinedSettings          bool          `json:"CRL Use CA Defined Settings"`
			AuthorityInformationAccessUsed   bool          `json:"Authority Information Access Used"`
			OCSPServiceLocatorURLIsCADefined bool          `json:"OCSP Service Locator URL Is CA Defined"`
			OCSPIssuerURLIsCADefined         bool          `json:"OCSP Issuer URL Is CA Defined"`
			CertificateExtensions            []string      `json:"Certificate Extensions"`
			CVCAccessRightsInspectionSystem  []interface{} `json:"CVC Access Rights (Inspection System)"`
			LDAPDNOrder                      bool          `json:"LDAP DN Order"`
			AvailableCAs                     []string      `json:"Available CAs"`
			AccountBindingNamespace          []interface{} `json:"Account Binding Namespace"`
		} `json:"solitude-workstation-msae"`
	} `json:"certificate-profiles"`
	PeerConnectors struct {
		SS01 struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			PeerEnabled                   bool   `json:"Peer Enabled"`
			URL                           string `json:"URL"`
			LongHangingConnectionsEnabled bool   `json:"Long Hanging Connections Enabled"`
			MinLongHangingConnections     int    `json:"Min Long Hanging Connections"`
			MaxLongHangingConnections     int    `json:"Max Long Hanging Connections"`
			AuthenticationKeyBinding      string `json:"Authentication Key Binding"`
		} `json:"SS-01"`
		RA01 struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			PeerEnabled                   bool   `json:"Peer Enabled"`
			URL                           string `json:"URL"`
			LongHangingConnectionsEnabled bool   `json:"Long Hanging Connections Enabled"`
			MinLongHangingConnections     int    `json:"Min Long Hanging Connections"`
			MaxLongHangingConnections     int    `json:"Max Long Hanging Connections"`
			AuthenticationKeyBinding      string `json:"Authentication Key Binding"`
		} `json:"RA-01"`
	} `json:"peer-connectors"`
	InternalKeyBindings struct {
		SS01 struct {
			ObjectType           string      `json:"Object Type"`
			Version              int         `json:"Version"`
			Name                 string      `json:"Name"`
			Type                 string      `json:"Type"`
			Status               string      `json:"Status"`
			CryptoToken          string      `json:"Crypto Token"`
			BoundCertificateSHA1 string      `json:"Bound Certificate SHA-1"`
			KeyPairAlias         string      `json:"Key Pair Alias"`
			NextKeyPairAlias     interface{} `json:"Next Key Pair Alias"`
			SignatureAlgorithm   string      `json:"Signature Algorithm"`
			TrustedCertificates  struct {
			} `json:"Trusted certificates"`
			SignedOnBehalfOfCAs struct {
			} `json:"Signed on behalf of CAs"`
			Properties struct {
				ProtocolAndCipherSuite string `json:"Protocol and Cipher Suite"`
			} `json:"Properties"`
		} `json:"SS-01"`
		LDAPAuthKeyBinding struct {
			ObjectType           string      `json:"Object Type"`
			Version              int         `json:"Version"`
			Name                 string      `json:"Name"`
			Type                 string      `json:"Type"`
			Status               string      `json:"Status"`
			CryptoToken          string      `json:"Crypto Token"`
			BoundCertificateSHA1 string      `json:"Bound Certificate SHA-1"`
			KeyPairAlias         string      `json:"Key Pair Alias"`
			NextKeyPairAlias     interface{} `json:"Next Key Pair Alias"`
			SignatureAlgorithm   string      `json:"Signature Algorithm"`
			TrustedCertificates  struct {
			} `json:"Trusted certificates"`
			SignedOnBehalfOfCAs struct {
			} `json:"Signed on behalf of CAs"`
			Properties struct {
				ProtocolAndCipherSuite string `json:"Protocol and Cipher Suite"`
			} `json:"Properties"`
		} `json:"LDAPAuthKeyBinding"`
		RA01 struct {
			ObjectType           string      `json:"Object Type"`
			Version              int         `json:"Version"`
			Name                 string      `json:"Name"`
			Type                 string      `json:"Type"`
			Status               string      `json:"Status"`
			CryptoToken          string      `json:"Crypto Token"`
			BoundCertificateSHA1 string      `json:"Bound Certificate SHA-1"`
			KeyPairAlias         string      `json:"Key Pair Alias"`
			NextKeyPairAlias     interface{} `json:"Next Key Pair Alias"`
			SignatureAlgorithm   string      `json:"Signature Algorithm"`
			TrustedCertificates  struct {
			} `json:"Trusted certificates"`
			SignedOnBehalfOfCAs struct {
			} `json:"Signed on behalf of CAs"`
			Properties struct {
				ProtocolAndCipherSuite string `json:"Protocol and Cipher Suite"`
			} `json:"Properties"`
		} `json:"RA-01"`
	} `json:"internal-key-bindings"`
	AcmeConfig struct {
		Itserver2D struct {
			ObjectType                                           string      `json:"Object Type"`
			Version                                              int         `json:"Version"`
			Name                                                 string      `json:"Name"`
			RANameGenerationScheme                               string      `json:"RA Name Generation Scheme"`
			RANameGenerationParams                               interface{} `json:"RA Name Generation Params"`
			EndEntityProfile                                     string      `json:"End Entity Profile"`
			PreAuthorizationAllowed                              bool        `json:"Pre-Authorization Allowed"`
			WildcardCertificateIssuanceAllowed                   bool        `json:"Wildcard Certificate Issuance Allowed"`
			WildcardCertificateIssuanceForHttp01ChallengeAllowed bool        `json:"Wildcard Certificate Issuance for http-01 Challenge Allowed"`
			DNSIdentifierChallengeTypes                          string      `json:"DNS Identifier Challenge Types"`
			TermsOfServiceRequireNewApproval                     bool        `json:"Terms of service require new approval"`
			AgreeToNewTermsOfServiceAllowed                      bool        `json:"Agree to new Terms of service allowed"`
			SiteURL                                              string      `json:"Site URL"`
			TermsOfServiceURL                                    string      `json:"Terms of Service URL"`
			TermsOfServiceChangeURL                              string      `json:"Terms of Service Change URL"`
			DNSResolver                                          string      `json:"DNS Resolver"`
			DNSPort                                              int         `json:"DNS Port"`
			DNSSECTrustAnchor                                    string      `json:"DNSSEC Trust Anchor"`
			ValidateDNSSEC                                       bool        `json:"Validate DNSSEC"`
			RequireClientAuthentication                          bool        `json:"Require Client Authentication"`
			RequireExternalAccountBinding                        bool        `json:"Require External Account Binding"`
			ExternalAccountBinding                               []struct {
				Version            int         `json:"Version"`
				TypeIdentifier     string      `json:"Type Identifier"`
				KeyIdentifier      string      `json:"Key Identifier"`
				MacAlgorithm       string      `json:"Mac Algorithm"`
				EncryptKey         bool        `json:"Encrypt Key"`
				EncryptionKeyId    string      `json:"Encryption Key Id"`
				EncryptionKeyAlias interface{} `json:"Encryption Key Alias"`
				SymmetricKey       string      `json:"Symmetric Key"`
			} `json:"External Account Binding"`
			RetryAfter                   int `json:"Retry After"`
			ApprovalProfileForNewAccount int `json:"Approval Profile for newAccount"`
			ApprovalProfileForKeyChange  int `json:"Approval Profile for keyChange"`
		} `json:"itserver2d"`
		Multiuse struct {
			ObjectType                                           string `json:"Object Type"`
			Version                                              int    `json:"Version"`
			Name                                                 string `json:"Name"`
			RANameGenerationScheme                               string `json:"RA Name Generation Scheme"`
			RANameGenerationParams                               string `json:"RA Name Generation Params"`
			EndEntityProfile                                     string `json:"End Entity Profile"`
			PreAuthorizationAllowed                              bool   `json:"Pre-Authorization Allowed"`
			WildcardCertificateIssuanceAllowed                   bool   `json:"Wildcard Certificate Issuance Allowed"`
			WildcardCertificateIssuanceForHttp01ChallengeAllowed bool   `json:"Wildcard Certificate Issuance for http-01 Challenge Allowed"`
			DNSIdentifierChallengeTypes                          string `json:"DNS Identifier Challenge Types"`
			TermsOfServiceRequireNewApproval                     bool   `json:"Terms of service require new approval"`
			AgreeToNewTermsOfServiceAllowed                      bool   `json:"Agree to new Terms of service allowed"`
			SiteURL                                              string `json:"Site URL"`
			TermsOfServiceURL                                    string `json:"Terms of Service URL"`
			TermsOfServiceChangeURL                              string `json:"Terms of Service Change URL"`
			DNSResolver                                          string `json:"DNS Resolver"`
			DNSPort                                              int    `json:"DNS Port"`
			DNSSECTrustAnchor                                    string `json:"DNSSEC Trust Anchor"`
			ValidateDNSSEC                                       bool   `json:"Validate DNSSEC"`
			RequireClientAuthentication                          bool   `json:"Require Client Authentication"`
			RequireExternalAccountBinding                        bool   `json:"Require External Account Binding"`
			RetryAfter                                           int    `json:"Retry After"`
			AuthorizedRedirectPorts                              string `json:"Authorized Redirect Ports"`
			ApprovalProfileForNewAccount                         int    `json:"Approval Profile for newAccount"`
			ApprovalProfileForKeyChange                          int    `json:"Approval Profile for keyChange"`
		} `json:"multiuse"`
		Acme30D struct {
			ObjectType                                           string `json:"Object Type"`
			Version                                              int    `json:"Version"`
			Name                                                 string `json:"Name"`
			RANameGenerationScheme                               string `json:"RA Name Generation Scheme"`
			RANameGenerationParams                               string `json:"RA Name Generation Params"`
			EndEntityProfile                                     string `json:"End Entity Profile"`
			PreAuthorizationAllowed                              bool   `json:"Pre-Authorization Allowed"`
			WildcardCertificateIssuanceAllowed                   bool   `json:"Wildcard Certificate Issuance Allowed"`
			WildcardCertificateIssuanceForHttp01ChallengeAllowed bool   `json:"Wildcard Certificate Issuance for http-01 Challenge Allowed"`
			DNSIdentifierChallengeTypes                          string `json:"DNS Identifier Challenge Types"`
			TermsOfServiceRequireNewApproval                     bool   `json:"Terms of service require new approval"`
			AgreeToNewTermsOfServiceAllowed                      bool   `json:"Agree to new Terms of service allowed"`
			SiteURL                                              string `json:"Site URL"`
			TermsOfServiceURL                                    string `json:"Terms of Service URL"`
			TermsOfServiceChangeURL                              string `json:"Terms of Service Change URL"`
			DNSResolver                                          string `json:"DNS Resolver"`
			DNSPort                                              int    `json:"DNS Port"`
			DNSSECTrustAnchor                                    string `json:"DNSSEC Trust Anchor"`
			ValidateDNSSEC                                       bool   `json:"Validate DNSSEC"`
			RequireClientAuthentication                          bool   `json:"Require Client Authentication"`
			RequireExternalAccountBinding                        bool   `json:"Require External Account Binding"`
			ExternalAccountBinding                               []struct {
				Version            int         `json:"Version"`
				TypeIdentifier     string      `json:"Type Identifier"`
				KeyIdentifier      string      `json:"Key Identifier"`
				MacAlgorithm       string      `json:"Mac Algorithm"`
				EncryptKey         bool        `json:"Encrypt Key"`
				EncryptionKeyId    interface{} `json:"Encryption Key Id"`
				EncryptionKeyAlias interface{} `json:"Encryption Key Alias"`
				SymmetricKey       string      `json:"Symmetric Key"`
			} `json:"External Account Binding"`
			RetryAfter                   int `json:"Retry After"`
			ApprovalProfileForNewAccount int `json:"Approval Profile for newAccount"`
			ApprovalProfileForKeyChange  int `json:"Approval Profile for keyChange"`
		} `json:"acme30d"`
		Solna struct {
			ObjectType                                           string `json:"Object Type"`
			Version                                              int    `json:"Version"`
			Name                                                 string `json:"Name"`
			RANameGenerationScheme                               string `json:"RA Name Generation Scheme"`
			RANameGenerationParams                               string `json:"RA Name Generation Params"`
			EndEntityProfile                                     string `json:"End Entity Profile"`
			PreAuthorizationAllowed                              bool   `json:"Pre-Authorization Allowed"`
			WildcardCertificateIssuanceAllowed                   bool   `json:"Wildcard Certificate Issuance Allowed"`
			WildcardCertificateIssuanceForHttp01ChallengeAllowed bool   `json:"Wildcard Certificate Issuance for http-01 Challenge Allowed"`
			DNSIdentifierChallengeTypes                          string `json:"DNS Identifier Challenge Types"`
			TermsOfServiceRequireNewApproval                     bool   `json:"Terms of service require new approval"`
			AgreeToNewTermsOfServiceAllowed                      bool   `json:"Agree to new Terms of service allowed"`
			SiteURL                                              string `json:"Site URL"`
			TermsOfServiceURL                                    string `json:"Terms of Service URL"`
			TermsOfServiceChangeURL                              string `json:"Terms of Service Change URL"`
			DNSResolver                                          string `json:"DNS Resolver"`
			DNSPort                                              int    `json:"DNS Port"`
			DNSSECTrustAnchor                                    string `json:"DNSSEC Trust Anchor"`
			ValidateDNSSEC                                       bool   `json:"Validate DNSSEC"`
			RequireClientAuthentication                          bool   `json:"Require Client Authentication"`
			RequireExternalAccountBinding                        bool   `json:"Require External Account Binding"`
			ExternalAccountBinding                               []struct {
				Version            int         `json:"Version"`
				TypeIdentifier     string      `json:"Type Identifier"`
				KeyIdentifier      string      `json:"Key Identifier"`
				MacAlgorithm       string      `json:"Mac Algorithm"`
				EncryptKey         bool        `json:"Encrypt Key"`
				EncryptionKeyId    string      `json:"Encryption Key Id"`
				EncryptionKeyAlias interface{} `json:"Encryption Key Alias"`
				SymmetricKey       string      `json:"Symmetric Key"`
			} `json:"External Account Binding"`
			RetryAfter                   int    `json:"Retry After"`
			AuthorizedRedirectPorts      string `json:"Authorized Redirect Ports"`
			ApprovalProfileForNewAccount int    `json:"Approval Profile for newAccount"`
			ApprovalProfileForKeyChange  int    `json:"Approval Profile for keyChange"`
		} `json:"solna"`
		GlobalAcmeConfiguration struct {
			ObjectType                        string `json:"Object Type"`
			Version                           int    `json:"Version"`
			DefaultAcmeConfiguration          string `json:"Default Acme Configuration"`
			ReplayNonceValidityInMilliseconds int    `json:"Replay-Nonce Validity In Milliseconds"`
		} `json:"global-acme-configuration"`
	} `json:"acme-config"`
	AdminRoles struct {
		SuperAdministratorRole struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string  `json:"Token Type"`
				Issuer      *string `json:"Issuer"`
				MatchWith   string  `json:"Match With"`
				MatchValue  string  `json:"Match Value"`
				Description string  `json:"Description,omitempty"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Field1              string `json:"/"`
				ProtocolCmp         string `json:"/protocol/cmp/"`
				ProtocolEst         string `json:"/protocol/est/"`
				ProtocolWebServices string `json:"/protocol/web_services/"`
				ProtocolAcme        string `json:"/protocol/acme/"`
				ProtocolRest        string `json:"/protocol/rest/"`
				ProtocolScep        string `json:"/protocol/scep/"`
			} `json:"Access Rules"`
		} `json:"Super Administrator Role"`
		CAAdmin struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                                               string `json:"/administrator/"`
				CaCarlSubG1                                                 string `json:"/ca/Carl-Sub-G1/"`
				CaCarlRootG1                                                string `json:"/ca/Carl-Root-G1/"`
				CaFunctionality                                             string `json:"/ca_functionality/"`
				CryptotokenView                                             string `json:"/cryptotoken/view/"`
				Endentityprofilesrules                                      string `json:"/endentityprofilesrules/"`
				InternalkeybindingView                                      string `json:"/internalkeybinding/view/"`
				RaFunctionality                                             string `json:"/ra_functionality/"`
				SecureauditAuditorSelect                                    string `json:"/secureaudit/auditor/select/"`
				ServicesView                                                string `json:"/services/view/"`
				SystemFunctionalityEditAdministratorPrivileges              string `json:"/system_functionality/edit_administrator_privileges/"`
				SystemFunctionalityViewAdministratorPrivileges              string `json:"/system_functionality/view_administrator_privileges/"`
				SystemFunctionalityViewAvailableCustomCertificateExtensions string `json:"/system_functionality/view_available_custom_certificate_extensions/"`
				SystemFunctionalityViewAvailableExtendedKeyUsages           string `json:"/system_functionality/view_available_extended_key_usages/"`
				SystemFunctionalityViewSystemconfiguration                  string `json:"/system_functionality/view_systemconfiguration/"`
				Validator                                                   string `json:"/validator/"`
			} `json:"Access Rules"`
		} `json:"CA-Admin"`
		RAVault struct {
			ObjectType  string        `json:"Object Type"`
			Version     int           `json:"Version"`
			Name        string        `json:"Name"`
			RoleMembers []interface{} `json:"Role Members"`
			Namespace   string        `json:"Namespace"`
			RAStyleId   int           `json:"RA Style Id"`
			AccessRules struct {
				Administrator                       string `json:"/administrator/"`
				CaITSubCA                           string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate    string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername          string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate      string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTlsServerIT   string `json:"/endentityprofilesrules/tlsServerIT/"`
				RaFunctionalityApproveEndEntity     string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity      string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity      string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity        string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityRevokeEndEntity      string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals        string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity        string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory string `json:"/ra_functionality/view_end_entity_history/"`
			} `json:"Access Rules"`
		} `json:"RA-Vault"`
		PeerSystems2 struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                        string `json:"/administrator/"`
				Ca                                   string `json:"/ca/"`
				CaFunctionalityCreateCertificate     string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId  string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername           string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCa                string `json:"/ca_functionality/view_ca/"`
				CaFunctionalityViewCertificate       string `json:"/ca_functionality/view_certificate/"`
				Endentityprofilesrules               string `json:"/endentityprofilesrules/"`
				Peerincoming                         string `json:"/peerincoming/"`
				ProtocolAcme                         string `json:"/protocol/acme/"`
				ProtocolCmp                          string `json:"/protocol/cmp/"`
				ProtocolEst                          string `json:"/protocol/est/"`
				ProtocolRest                         string `json:"/protocol/rest/"`
				ProtocolScep                         string `json:"/protocol/scep/"`
				ProtocolWebServices                  string `json:"/protocol/web_services/"`
				RaFunctionalityApproveEndEntity      string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity       string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity       string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity         string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery           string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity       string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals         string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity         string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory  string `json:"/ra_functionality/view_end_entity_history/"`
				RaFunctionalityViewEndEntityProfiles string `json:"/ra_functionality/view_end_entity_profiles/"`
				RaMasterInvokeApi                    string `json:"/ra_master/invoke_api/"`
			} `json:"Access Rules"`
		} `json:"PeerSystems2"`
		RA struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType  string `json:"Token Type"`
				Issuer     string `json:"Issuer"`
				MatchWith  string `json:"Match With"`
				MatchValue string `json:"Match Value"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                       string `json:"/administrator/"`
				Ca                                  string `json:"/ca/"`
				CaFunctionalityCreateCertificate    string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername          string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate      string `json:"/ca_functionality/view_certificate/"`
				Endentityprofilesrules              string `json:"/endentityprofilesrules/"`
				RaFunctionalityApproveEndEntity     string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity      string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity      string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity        string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery          string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity      string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals        string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity        string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect            string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA"`
		RAMultiuse struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                        string `json:"/administrator/"`
				CaITSubCA                            string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate     string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId  string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername           string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate       string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTestMultiIssue string `json:"/endentityprofilesrules/testMultiIssue/"`
				RaFunctionalityApproveEndEntity      string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity       string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity       string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity         string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery           string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity       string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals         string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity         string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory  string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect             string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-multiuse"`
		RAWorkspc1 struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                       string `json:"/administrator/"`
				CaFunctionalityCreateCertificate    string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername          string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate      string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesIntuneDevice  string `json:"/endentityprofilesrules/intuneDevice/"`
				RaFunctionalityApproveEndEntity     string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity      string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity      string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity        string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery          string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity      string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals        string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity        string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory string `json:"/ra_functionality/view_end_entity_history/"`
			} `json:"Access Rules"`
		} `json:"RA-Workspc1"`
		Auditor struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType  string `json:"Token Type"`
				Issuer     string `json:"Issuer"`
				MatchWith  string `json:"Match With"`
				MatchValue string `json:"Match Value"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                                               string `json:"/administrator/"`
				Ca                                                          string `json:"/ca/"`
				CaFunctionalityViewApprovalProfiles                         string `json:"/ca_functionality/view_approval_profiles/"`
				CaFunctionalityViewCa                                       string `json:"/ca_functionality/view_ca/"`
				CaFunctionalityViewCertificate                              string `json:"/ca_functionality/view_certificate/"`
				CaFunctionalityViewCertificateProfiles                      string `json:"/ca_functionality/view_certificate_profiles/"`
				CaFunctionalityViewPublisher                                string `json:"/ca_functionality/view_publisher/"`
				CaFunctionalityViewValidator                                string `json:"/ca_functionality/view_validator/"`
				CryptotokenView                                             string `json:"/cryptotoken/view/"`
				Endentityprofilesrules                                      string `json:"/endentityprofilesrules/"`
				InternalkeybindingView                                      string `json:"/internalkeybinding/view/"`
				PeerView                                                    string `json:"/peer/view/"`
				RaFunctionalityViewEndEntity                                string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory                         string `json:"/ra_functionality/view_end_entity_history/"`
				RaFunctionalityViewEndEntityProfiles                        string `json:"/ra_functionality/view_end_entity_profiles/"`
				SecureauditAuditorSelect                                    string `json:"/secureaudit/auditor/select/"`
				ServicesView                                                string `json:"/services/view/"`
				SystemFunctionalityViewAdministratorPrivileges              string `json:"/system_functionality/view_administrator_privileges/"`
				SystemFunctionalityViewAvailableCustomCertificateExtensions string `json:"/system_functionality/view_available_custom_certificate_extensions/"`
				SystemFunctionalityViewAvailableExtendedKeyUsages           string `json:"/system_functionality/view_available_extended_key_usages/"`
				SystemFunctionalityViewSystemconfiguration                  string `json:"/system_functionality/view_systemconfiguration/"`
				Validator                                                   string `json:"/validator/"`
			} `json:"Access Rules"`
		} `json:"Auditor"`
		RAAnsible struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                       string `json:"/administrator/"`
				CaITSubCA                           string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate    string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername          string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate      string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTlsServerIT   string `json:"/endentityprofilesrules/tlsServerIT/"`
				RaFunctionalityApproveEndEntity     string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity      string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity      string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity        string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery          string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity      string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals        string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity        string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect            string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-ansible"`
		RAServiceNow struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType  string `json:"Token Type"`
				Issuer     string `json:"Issuer"`
				MatchWith  string `json:"Match With"`
				MatchValue string `json:"Match Value"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                                  string `json:"/administrator/"`
				Ca                                             string `json:"/ca/"`
				CaFunctionality                                string `json:"/ca_functionality/"`
				CryptotokenView                                string `json:"/cryptotoken/view/"`
				Endentityprofilesrules                         string `json:"/endentityprofilesrules/"`
				InternalkeybindingDelete                       string `json:"/internalkeybinding/delete/"`
				InternalkeybindingModify                       string `json:"/internalkeybinding/modify/"`
				InternalkeybindingView                         string `json:"/internalkeybinding/view/"`
				RaFunctionality                                string `json:"/ra_functionality/"`
				SecureauditAuditorSelect                       string `json:"/secureaudit/auditor/select/"`
				SystemFunctionalityEditAdministratorPrivileges string `json:"/system_functionality/edit_administrator_privileges/"`
				SystemFunctionalityViewAdministratorPrivileges string `json:"/system_functionality/view_administrator_privileges/"`
				Validator                                      string `json:"/validator/"`
			} `json:"Access Rules"`
		} `json:"RA-ServiceNow"`
		MgmtApproval struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                                    string `json:"/administrator/"`
				CaITSubCA                                        string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate                 string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId              string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername                       string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate                   string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesUserAuthenticationApproval string `json:"/endentityprofilesrules/userAuthenticationApproval/"`
				RaFunctionalityApproveEndEntity                  string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityViewApprovals                     string `json:"/ra_functionality/view_approvals/"`
			} `json:"Access Rules"`
		} `json:"Mgmt-Approval"`
		RAEST struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType  string `json:"Token Type"`
				Issuer     string `json:"Issuer"`
				MatchWith  string `json:"Match With"`
				MatchValue string `json:"Match Value"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                           string `json:"/administrator/"`
				CaITSubCA                               string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate        string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId     string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername              string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate          string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesVpnFirewall       string `json:"/endentityprofilesrules/vpnFirewall/"`
				EndentityprofilesrulesVpnFirewallTest1D string `json:"/endentityprofilesrules/vpnFirewallTest1d/"`
				EndentityprofilesrulesEstDevice         string `json:"/endentityprofilesrules/estDevice/"`
				RaFunctionalityCreateEndEntity          string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity          string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity            string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityViewApprovals            string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity            string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory     string `json:"/ra_functionality/view_end_entity_history/"`
			} `json:"Access Rules"`
		} `json:"RA-EST"`
		RADPoDCompany2 struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                           string `json:"/administrator/"`
				CaDPodSubCA                             string `json:"/ca/DPod-SubCA/"`
				CaFunctionalityCreateCertificate        string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId     string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername              string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate          string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTlsServerCompany2 string `json:"/endentityprofilesrules/tlsServerCompany2/"`
				RaFunctionalityApproveEndEntity         string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity          string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity          string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity            string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityRevokeEndEntity          string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals            string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity            string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory     string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect                string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-DPoD-Company2"`
		UserEnrollment struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string      `json:"Token Type"`
				Issuer      interface{} `json:"Issuer"`
				MatchWith   string      `json:"Match With"`
				MatchValue  string      `json:"Match Value"`
				Description string      `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				CaITSubCA                                                     string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate                              string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId                           string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername                                    string `json:"/ca_functionality/use_username/"`
				EndentityprofilesrulesTlsServerApprovalITCreateEndEntity      string `json:"/endentityprofilesrules/tlsServerApprovalIT/create_end_entity/"`
				EndentityprofilesrulesTlsServerApprovalITDeleteEndEntity      string `json:"/endentityprofilesrules/tlsServerApprovalIT/delete_end_entity/"`
				EndentityprofilesrulesTlsServerApprovalITViewEndEntity        string `json:"/endentityprofilesrules/tlsServerApprovalIT/view_end_entity/"`
				EndentityprofilesrulesTlsServerApprovalITViewEndEntityHistory string `json:"/endentityprofilesrules/tlsServerApprovalIT/view_end_entity_history/"`
				RaFunctionalityCreateEndEntity                                string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity                                string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityViewEndEntity                                  string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory                           string `json:"/ra_functionality/view_end_entity_history/"`
				RaFunctionalityViewEndEntityProfiles                          string `json:"/ra_functionality/view_end_entity_profiles/"`
			} `json:"Access Rules"`
		} `json:"UserEnrollment"`
		RASolna struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType  string `json:"Token Type"`
				Issuer     string `json:"Issuer"`
				MatchWith  string `json:"Match With"`
				MatchValue string `json:"Match Value"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                              string `json:"/administrator/"`
				CaITSubCA                                  string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate           string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId        string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername                 string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate             string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesSolnaWorkstation     string `json:"/endentityprofilesrules/solna-workstation/"`
				EndentityprofilesrulesSolnaTlsServerAcme2D string `json:"/endentityprofilesrules/solna-tlsServerAcme2d/"`
				RaFunctionalityApproveEndEntity            string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity             string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity             string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity               string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery                 string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity             string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals               string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity               string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory        string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect                   string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-Solna"`
		TestAuditor struct {
			ObjectType  string        `json:"Object Type"`
			Version     int           `json:"Version"`
			Name        string        `json:"Name"`
			RoleMembers []interface{} `json:"Role Members"`
			Namespace   string        `json:"Namespace"`
			RAStyleId   int           `json:"RA Style Id"`
			AccessRules struct {
				Administrator                       string `json:"/administrator/"`
				CaITSubCA                           string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate    string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername          string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate      string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesScepDevice    string `json:"/endentityprofilesrules/scepDevice/"`
				RaFunctionalityApproveEndEntity     string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity      string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityEditEndEntity        string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityViewApprovals        string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity        string `json:"/ra_functionality/view_end_entity/"`
				SecureauditAuditorSelect            string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"test-auditor"`
		RAPeering struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                               string `json:"/administrator/"`
				CaManagementCA                              string `json:"/ca/ManagementCA/"`
				CaKeyVaultSubCA                             string `json:"/ca/KeyVault-Sub-CA/"`
				CaITSubCA                                   string `json:"/ca/IT-Sub-CA/"`
				CaDPodSubCA                                 string `json:"/ca/DPod-SubCA/"`
				CaFunctionalityCreateCertificate            string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId         string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername                  string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCa                       string `json:"/ca_functionality/view_ca/"`
				CaFunctionalityViewCertificate              string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTlsServerAcme60D      string `json:"/endentityprofilesrules/tlsServerAcme60d/"`
				EndentityprofilesrulesVpnFirewall10M        string `json:"/endentityprofilesrules/vpnFirewall10m/"`
				EndentityprofilesrulesDomainController      string `json:"/endentityprofilesrules/domainController/"`
				EndentityprofilesrulesAdminInternal         string `json:"/endentityprofilesrules/AdminInternal/"`
				EndentityprofilesrulesSignersIT             string `json:"/endentityprofilesrules/signersIT/"`
				EndentityprofilesrulesScepDevice            string `json:"/endentityprofilesrules/scepDevice/"`
				EndentityprofilesrulesVpnFirewall           string `json:"/endentityprofilesrules/vpnFirewall/"`
				EndentityprofilesrulesIntuneDevice          string `json:"/endentityprofilesrules/intuneDevice/"`
				EndentityprofilesrulesTlsServerMgmt         string `json:"/endentityprofilesrules/tlsServerMgmt/"`
				EndentityprofilesrulesUserAuthentication    string `json:"/endentityprofilesrules/userAuthentication/"`
				EndentityprofilesrulesUserDigitalSignature  string `json:"/endentityprofilesrules/userDigitalSignature/"`
				EndentityprofilesrulesRemoteDesktopAuth     string `json:"/endentityprofilesrules/remoteDesktopAuth/"`
				EndentityprofilesrulesTlsServerAcme2D       string `json:"/endentityprofilesrules/tlsServerAcme2d/"`
				EndentityprofilesrulesTlsServerAcme30D      string `json:"/endentityprofilesrules/tlsServerAcme30d/"`
				EndentityprofilesrulesTlsServerAcme         string `json:"/endentityprofilesrules/tlsServerAcme/"`
				EndentityprofilesrulesTlsPeerConnMgmt       string `json:"/endentityprofilesrules/tlsPeerConnMgmt/"`
				EndentityprofilesrulesWorkstations          string `json:"/endentityprofilesrules/workstations/"`
				EndentityprofilesrulesTimeStampingAuthority string `json:"/endentityprofilesrules/timeStampingAuthority/"`
				EndentityprofilesrulesOcspSigner            string `json:"/endentityprofilesrules/ocspSigner/"`
				EndentityprofilesrulesTlsServerIT           string `json:"/endentityprofilesrules/tlsServerIT/"`
				EndentityprofilesrulesEstDevice             string `json:"/endentityprofilesrules/estDevice/"`
				EndentityprofilesrulesTlsServerApprovalIT   string `json:"/endentityprofilesrules/tlsServerApprovalIT/"`
				EndentityprofilesrulesUserEncryption        string `json:"/endentityprofilesrules/userEncryption/"`
				ProtocolAcme                                string `json:"/protocol/acme/"`
				ProtocolCmp                                 string `json:"/protocol/cmp/"`
				ProtocolEst                                 string `json:"/protocol/est/"`
				ProtocolRest                                string `json:"/protocol/rest/"`
				ProtocolScep                                string `json:"/protocol/scep/"`
				ProtocolWebServices                         string `json:"/protocol/web_services/"`
				RaFunctionalityApproveEndEntity             string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity              string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity              string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity                string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityRevokeEndEntity              string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals                string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity                string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory         string `json:"/ra_functionality/view_end_entity_history/"`
				RaFunctionalityViewEndEntityProfiles        string `json:"/ra_functionality/view_end_entity_profiles/"`
				RaMasterInvokeApi                           string `json:"/ra_master/invoke_api/"`
			} `json:"Access Rules"`
		} `json:"RA-Peering"`
		RAApproval struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                                    string `json:"/administrator/"`
				CaITSubCA                                        string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate                 string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId              string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername                       string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate                   string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesUserAuthenticationApproval string `json:"/endentityprofilesrules/userAuthenticationApproval/"`
				RaFunctionalityApproveEndEntity                  string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity                   string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity                   string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity                     string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityRevokeEndEntity                   string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals                     string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity                     string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory              string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect                         string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-Approval"`
		ItSs01PkihostedDevC2CompanyComRA struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string      `json:"Token Type"`
				Issuer      interface{} `json:"Issuer"`
				MatchWith   string      `json:"Match With"`
				MatchValue  string      `json:"Match Value"`
				Description string      `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Ca                                   string `json:"/ca/"`
				CaFunctionalityCreateCertificate     string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId  string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername           string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate       string `json:"/ca_functionality/view_certificate/"`
				Endentityprofilesrules               string `json:"/endentityprofilesrules/"`
				RaFunctionalityApproveEndEntity      string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity       string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity       string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity         string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityViewApprovals         string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity         string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory  string `json:"/ra_functionality/view_end_entity_history/"`
				RaFunctionalityViewEndEntityProfiles string `json:"/ra_functionality/view_end_entity_profiles/"`
				RaMasterInvokeApi                    string `json:"/ra_master/invoke_api/"`
			} `json:"Access Rules"`
		} `json:"it-ss01.pkihosted-dev.c2company.com RA"`
		RACleveland struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType  string `json:"Token Type"`
				Issuer     string `json:"Issuer"`
				MatchWith  string `json:"Match With"`
				MatchValue string `json:"Match Value"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                                string `json:"/administrator/"`
				CaDPodSubCA                                  string `json:"/ca/DPod-SubCA/"`
				CaFunctionalityCreateCertificate             string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId          string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername                   string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate               string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesClevelandWorkstation2D string `json:"/endentityprofilesrules/cleveland-workstation2d/"`
				RaFunctionalityApproveEndEntity              string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity               string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity               string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity                 string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityKeyrecovery                   string `json:"/ra_functionality/keyrecovery/"`
				RaFunctionalityRevokeEndEntity               string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals                 string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity                 string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory          string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect                     string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-Cleveland"`
		RADPoDCompany1 struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                           string `json:"/administrator/"`
				CaDPodSubCA                             string `json:"/ca/DPod-SubCA/"`
				CaFunctionalityCreateCertificate        string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId     string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername              string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate          string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTlsServerCompany1 string `json:"/endentityprofilesrules/tlsServerCompany1/"`
				RaFunctionalityApproveEndEntity         string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity          string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity          string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity            string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityRevokeEndEntity          string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals            string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity            string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory     string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect                string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-DPoD-Company1"`
		RALemur struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			RoleMembers []struct {
				TokenType   string `json:"Token Type"`
				Issuer      string `json:"Issuer"`
				MatchWith   string `json:"Match With"`
				MatchValue  string `json:"Match Value"`
				Description string `json:"Description"`
			} `json:"Role Members"`
			Namespace   string `json:"Namespace"`
			RAStyleId   int    `json:"RA Style Id"`
			AccessRules struct {
				Administrator                            string `json:"/administrator/"`
				CaITSubCA                                string `json:"/ca/IT-Sub-CA/"`
				CaFunctionalityCreateCertificate         string `json:"/ca_functionality/create_certificate/"`
				CaFunctionalityUseApprovalRequestId      string `json:"/ca_functionality/use_approval_request_id/"`
				CaFunctionalityUseUsername               string `json:"/ca_functionality/use_username/"`
				CaFunctionalityViewCertificate           string `json:"/ca_functionality/view_certificate/"`
				EndentityprofilesrulesTlsServerAcme60D   string `json:"/endentityprofilesrules/tlsServerAcme60d/"`
				EndentityprofilesrulesIntuneDevice       string `json:"/endentityprofilesrules/intuneDevice/"`
				EndentityprofilesrulesUserAuthentication string `json:"/endentityprofilesrules/userAuthentication/"`
				EndentityprofilesrulesTlsServerAcme2D    string `json:"/endentityprofilesrules/tlsServerAcme2d/"`
				EndentityprofilesrulesTlsServerAcme30D   string `json:"/endentityprofilesrules/tlsServerAcme30d/"`
				EndentityprofilesrulesTlsServerIT        string `json:"/endentityprofilesrules/tlsServerIT/"`
				RaFunctionalityApproveEndEntity          string `json:"/ra_functionality/approve_end_entity/"`
				RaFunctionalityCreateEndEntity           string `json:"/ra_functionality/create_end_entity/"`
				RaFunctionalityDeleteEndEntity           string `json:"/ra_functionality/delete_end_entity/"`
				RaFunctionalityEditEndEntity             string `json:"/ra_functionality/edit_end_entity/"`
				RaFunctionalityRevokeEndEntity           string `json:"/ra_functionality/revoke_end_entity/"`
				RaFunctionalityViewApprovals             string `json:"/ra_functionality/view_approvals/"`
				RaFunctionalityViewEndEntity             string `json:"/ra_functionality/view_end_entity/"`
				RaFunctionalityViewEndEntityHistory      string `json:"/ra_functionality/view_end_entity_history/"`
				SecureauditAuditorSelect                 string `json:"/secureaudit/auditor/select/"`
			} `json:"Access Rules"`
		} `json:"RA-Lemur"`
	} `json:"admin-roles"`
	CryptoTokens struct {
		Peering struct {
			ObjectType          string `json:"Object Type"`
			Version             int    `json:"Version"`
			Name                string `json:"Name"`
			Used                bool   `json:"Used"`
			PKCS11Library       string `json:"PKCS11 Library"`
			PKCS11ReferenceType string `json:"PKCS11 Reference Type"`
			PKCS11Reference     string `json:"PKCS11 Reference"`
			PKCS11AttributeFile string `json:"PKCS11 Attribute File"`
			KeyPairInfo         []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"Peering"`
		DPodSubCA struct {
			ObjectType          string `json:"Object Type"`
			Version             int    `json:"Version"`
			Name                string `json:"Name"`
			Used                bool   `json:"Used"`
			PKCS11Library       string `json:"PKCS11 Library"`
			PKCS11ReferenceType string `json:"PKCS11 Reference Type"`
			PKCS11Reference     string `json:"PKCS11 Reference"`
			PKCS11AttributeFile string `json:"PKCS11 Attribute File"`
			KeyPairInfo         []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"DPod-SubCA"`
		CarlPki struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			Used        bool   `json:"Used"`
			KeyPairInfo []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"CarlPki"`
		ITSubCA struct {
			ObjectType          string `json:"Object Type"`
			Version             int    `json:"Version"`
			Name                string `json:"Name"`
			Used                bool   `json:"Used"`
			PKCS11Library       string `json:"PKCS11 Library"`
			PKCS11ReferenceType string `json:"PKCS11 Reference Type"`
			PKCS11Reference     string `json:"PKCS11 Reference"`
			PKCS11AttributeFile string `json:"PKCS11 Attribute File"`
			KeyPairInfo         []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"IT-Sub-CA"`
		Demotoken struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			Used        bool   `json:"Used"`
			KeyPairInfo []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"demotoken"`
		ManagementCA struct {
			ObjectType          string `json:"Object Type"`
			Version             int    `json:"Version"`
			Name                string `json:"Name"`
			Used                bool   `json:"Used"`
			PKCS11Library       string `json:"PKCS11 Library"`
			PKCS11ReferenceType string `json:"PKCS11 Reference Type"`
			PKCS11Reference     string `json:"PKCS11 Reference"`
			PKCS11AttributeFile string `json:"PKCS11 Attribute File"`
			KeyPairInfo         []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"ManagementCA"`
		KeyVault struct {
			ObjectType                 string        `json:"Object Type"`
			Version                    int           `json:"Version"`
			Name                       string        `json:"Name"`
			Used                       bool          `json:"Used"`
			KeyVaultType               string        `json:"Key Vault Type"`
			KeyVaultName               string        `json:"Key Vault Name"`
			KeyVaultClientID           string        `json:"Key Vault ClientID"`
			KeyVaultAuthenticationType string        `json:"Key Vault Authentication Type"`
			KeyVaultKeyBinding         interface{}   `json:"Key Vault Key Binding"`
			KeyPairInfo                []interface{} `json:"Key Pair Info"`
			AuthenticationCode         string        `json:"Authentication Code"`
			Type                       string        `json:"Type"`
			Active                     bool          `json:"Active"`
		} `json:"KeyVault"`
		TestCps struct {
			ObjectType         string        `json:"Object Type"`
			Version            int           `json:"Version"`
			Name               string        `json:"Name"`
			Used               bool          `json:"Used"`
			KeyPairInfo        []interface{} `json:"Key Pair Info"`
			AuthenticationCode string        `json:"Authentication Code"`
			Type               string        `json:"Type"`
			Active             bool          `json:"Active"`
		} `json:"test-cps"`
		ITRootCA struct {
			ObjectType          string `json:"Object Type"`
			Version             int    `json:"Version"`
			Name                string `json:"Name"`
			Used                bool   `json:"Used"`
			PKCS11Library       string `json:"PKCS11 Library"`
			PKCS11ReferenceType string `json:"PKCS11 Reference Type"`
			PKCS11Reference     string `json:"PKCS11 Reference"`
			PKCS11AttributeFile string `json:"PKCS11 Attribute File"`
			KeyPairInfo         []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"IT-Root-CA"`
		SSHCA01 struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			Used        bool   `json:"Used"`
			KeyPairInfo []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"SSH-CA-01"`
		ActiveDirectory struct {
			ObjectType  string `json:"Object Type"`
			Version     int    `json:"Version"`
			Name        string `json:"Name"`
			Used        bool   `json:"Used"`
			KeyPairInfo []struct {
				Alias                string `json:"Alias"`
				KeyAlgorithm         string `json:"Key Algorithm"`
				KeySpecification     string `json:"Key Specification"`
				SubjectKeyIdentifier string `json:"Subject Key Identifier"`
			} `json:"Key Pair Info"`
			AuthenticationCode string `json:"Authentication Code"`
			Type               string `json:"Type"`
			Active             bool   `json:"Active"`
			AutoActivation     bool   `json:"Auto Activation"`
		} `json:"active-directory"`
	} `json:"crypto-tokens"`
	Publishers struct {
		CTLogs struct {
			ObjectType string `json:"Object Type"`
			Version    int    `json:"Version"`
			Name       string `json:"Name"`
			Type       string `json:"Type"`
			Settings   struct {
			} `json:"Settings"`
			OnlyUseInQueue           bool   `json:"Only Use in Queue"`
			UseSafeDirectPublishing  bool   `json:"Use Safe Direct Publishing"`
			KeepPublishedInQueue     bool   `json:"Keep Published in Queue"`
			UseQueueForCRLs          bool   `json:"Use Queue for CRLs"`
			UseQueueForCertificates  bool   `json:"Use Queue for Certificates"`
			UseQueueForOCSPResponses bool   `json:"Use Queue for OCSP Responses"`
			Description              string `json:"Description"`
		} `json:"CT-logs"`
		AzureBlobCA struct {
			ObjectType string `json:"Object Type"`
			Version    int    `json:"Version"`
			Name       string `json:"Name"`
			Type       string `json:"Type"`
			Settings   struct {
				PeerThatWillConnectToAzureAndAuthenticationMethod string `json:"Peer that will connect to Azure and Authentication Method"`
				AzureStorageResourceName                          string `json:"Azure Storage Resource Name"`
				FormatOfCertificateSAzureBlobName                 string `json:"Format of Certificate's Azure Blob Name"`
				AzureBlobStorageAccount                           string `json:"Azure Blob Storage Account"`
				CertificateFormat                                 string `json:"Certificate Format"`
				FormatOfCRLSAzureBlobName                         string `json:"Format of CRL's Azure Blob Name"`
				AzureBlobContainer                                string `json:"Azure Blob Container"`
				CrlFormat                                         string `json:"Crl Format"`
				AzureStorageBlobEndpoint                          string `json:"Azure Storage Blob Endpoint"`
				MicrosoftTenantID                                 string `json:"Microsoft Tenant ID"`
				MicrosoftApplicationID                            string `json:"Microsoft Application ID"`
				MicrosoftAuthenticationSecret                     string `json:"Microsoft Authentication Secret"`
			} `json:"Settings"`
			OnlyUseInQueue           bool   `json:"Only Use in Queue"`
			UseSafeDirectPublishing  bool   `json:"Use Safe Direct Publishing"`
			KeepPublishedInQueue     bool   `json:"Keep Published in Queue"`
			UseQueueForCRLs          bool   `json:"Use Queue for CRLs"`
			UseQueueForCertificates  bool   `json:"Use Queue for Certificates"`
			UseQueueForOCSPResponses bool   `json:"Use Queue for OCSP Responses"`
			Description              string `json:"Description"`
		} `json:"Azure-Blob-CA"`
		Ad struct {
			ObjectType                           string        `json:"Object Type"`
			Version                              int           `json:"Version"`
			Name                                 string        `json:"Name"`
			Type                                 string        `json:"Type"`
			Hostnames                            string        `json:"Hostnames"`
			Port                                 string        `json:"Port"`
			BaseDN                               string        `json:"Base DN"`
			LoginDN                              string        `json:"Login DN"`
			LoginPassword                        string        `json:"Login Password"`
			ConnectionTimeout                    int           `json:"Connection Timeout"`
			ReadTimeout                          int           `json:"Read Timeout"`
			StoreTimeout                         int           `json:"Store Timeout"`
			CreateNonExistingUsers               bool          `json:"Create Non-Existing Users"`
			ModifyExistingUsers                  bool          `json:"Modify Existing Users"`
			OverwriteExistingAttributes          bool          `json:"Overwrite Existing Attributes"`
			AddNonExistingAttributes             bool          `json:"Add Non-Existing Attributes"`
			CreateIntermediateNodes              bool          `json:"Create Intermediate Nodes"`
			AddMultipleCertificatePerUser        bool          `json:"Add Multiple Certificate per User"`
			RemoveCertificatesWhenRevoked        bool          `json:"Remove Certificates When Revoked"`
			RemoveLDAPUserWhenCertificateRevoked bool          `json:"Remove LDAP User When Certificate Revoked"`
			SetUserPasswordAttribute             bool          `json:"Set userPassword Attribute"`
			UserObjectClass                      string        `json:"User Object Class"`
			CAObjectClass                        string        `json:"CA Object Class"`
			UserCertificateAttribute             string        `json:"User Certificate Attribute"`
			CACertificateAttribute               string        `json:"CA Certificate Attribute"`
			CRLAttribute                         string        `json:"CRL Attribute"`
			DeltaCRLAttribute                    string        `json:"Delta CRL Attribute"`
			ARLAttribute                         string        `json:"ARL Attribute"`
			LDAPLocationFieldsFromCertificateDN  []interface{} `json:"LDAP Location Fields from Certificate DN"`
			UseCustomDNOrder                     bool          `json:"Use Custom DN Order"`
			OnlyUseInQueue                       bool          `json:"Only Use in Queue"`
			UseSafeDirectPublishing              bool          `json:"Use Safe Direct Publishing"`
			KeepPublishedInQueue                 bool          `json:"Keep Published in Queue"`
			UseQueueForCRLs                      bool          `json:"Use Queue for CRLs"`
			UseQueueForCertificates              bool          `json:"Use Queue for Certificates"`
			UseQueueForOCSPResponses             bool          `json:"Use Queue for OCSP Responses"`
			Description                          string        `json:"Description"`
		} `json:"ad"`
		AmazonS3 struct {
			ObjectType string `json:"Object Type"`
			Version    int    `json:"Version"`
			Name       string `json:"Name"`
			Type       string `json:"Type"`
			Settings   struct {
				S3CrlFileFormat         string `json:"s3.crl.file.format"`
				S3CertStoreSeparatePath string `json:"s3.cert.store.separate.path"`
				S3CrlKeyPrefix          string `json:"s3.crl.key.prefix"`
				S3CertFileFormat        string `json:"s3.cert.file.format"`
				S3CrlBucketName         string `json:"s3.crl.bucket.name"`
				S3CrlFileNameFormat     string `json:"s3.crl.file.name.format"`
				S3CertFileNameFormat    string `json:"s3.cert.file.name.format"`
				S3CertKeyPrefix         string `json:"s3.cert.key.prefix"`
				S3CertBucketName        string `json:"s3.cert.bucket.name"`
			} `json:"Settings"`
			OnlyUseInQueue           bool   `json:"Only Use in Queue"`
			UseSafeDirectPublishing  bool   `json:"Use Safe Direct Publishing"`
			KeepPublishedInQueue     bool   `json:"Keep Published in Queue"`
			UseQueueForCRLs          bool   `json:"Use Queue for CRLs"`
			UseQueueForCertificates  bool   `json:"Use Queue for Certificates"`
			UseQueueForOCSPResponses bool   `json:"Use Queue for OCSP Responses"`
			Description              string `json:"Description"`
		} `json:"AmazonS3"`
	} `json:"publishers"`
	OcspConfiguration struct {
		OcspConfiguration struct {
			ObjectType                                      string `json:"Object Type"`
			Version                                         int    `json:"Version"`
			Name                                            string `json:"Name"`
			DefaultResponder                                string `json:"Default Responder"`
			ResponderIDTypeForCAs                           string `json:"Responder ID Type for CAs"`
			NonceExtensionInOCSPRepliesFromCAsEnabled       bool   `json:"Nonce extension in OCSP replies from CAs enabled"`
			OCSPSigningCacheUpdateEnabled                   bool   `json:"OCSP signing cache update enabled"`
			OCSPExplicitNoCacheUnauthorizedResponsesEnabled bool   `json:"OCSP explicit no cache unauthorized responses enabled"`
			AuditLoggingEnabled                             bool   `json:"Audit Logging Enabled"`
			AuditLogPattern                                 string `json:"Audit Log Pattern"`
			AuditLogValues                                  string `json:"Audit Log Values"`
			TransactionLoggingEnabled                       bool   `json:"Transaction Logging Enabled"`
			TransactionLogPattern                           string `json:"Transaction Log Pattern"`
			TransactionLogValues                            string `json:"Transaction Log Values"`
			OCSPLoggingDateFormat                           string `json:"OCSP Logging Date Format"`
		} `json:"ocsp-configuration"`
	} `json:"ocsp-configuration"`
	CertificationAuthorities struct {
		TestCps struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			FinishUser                             bool        `json:"Finish User"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"test-cps"`
		CarlRootG1 struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			SubjectAlternativeName                 string      `json:"Subject Alternative Name"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			CAIssuerURI                            []string    `json:"CA Issuer URI"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			FinishUser                             bool        `json:"Finish User"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"Carl-Root-G1"`
		ManagementCA struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			Description                   string `json:"Description"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			CAIssuerURI                            []string    `json:"CA Issuer URI"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			OCSPServiceDefaultURI                  string      `json:"OCSP Service Default URI"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			FinishUser                             bool        `json:"Finish User"`
			CAHealthcheckEnabled                   bool        `json:"CA Healthcheck Enabled"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"ManagementCA"`
		KeyVaultSubCA struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			SubjectAlternativeName                 string      `json:"Subject Alternative Name"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			OCSPServiceDefaultURI                  string      `json:"OCSP Service Default URI"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			CRLPublishers                          []string    `json:"CRL Publishers"`
			FinishUser                             bool        `json:"Finish User"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"KeyVault-Sub-CA"`
		ITSubCA struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			Description                   string `json:"Description"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			SubjectAlternativeName                 string      `json:"Subject Alternative Name"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			CAIssuerURI                            []string    `json:"CA Issuer URI"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			OCSPServiceDefaultURI                  string      `json:"OCSP Service Default URI"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			Validators                             []string    `json:"Validators"`
			FinishUser                             bool        `json:"Finish User"`
			CAHealthcheckEnabled                   bool        `json:"CA Healthcheck Enabled"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"IT-Sub-CA"`
		DPodSubCA struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			SubjectAlternativeName                 interface{} `json:"Subject Alternative Name"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			CAIssuerURI                            []string    `json:"CA Issuer URI"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			OCSPServiceDefaultURI                  string      `json:"OCSP Service Default URI"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			FinishUser                             bool        `json:"Finish User"`
			CAHealthcheckEnabled                   bool        `json:"CA Healthcheck Enabled"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"DPod-SubCA"`
		ITRootCA struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			Description                   string `json:"Description"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			SubjectAlternativeName                 string      `json:"Subject Alternative Name"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			PrintableStringEncodingInDN            bool        `json:"Printable String Encoding in DN"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			CAIssuerURI                            []string    `json:"CA Issuer URI"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			OCSPServiceDefaultURI                  string      `json:"OCSP Service Default URI"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			CRLPublishers                          []string    `json:"CRL Publishers"`
			FinishUser                             bool        `json:"Finish User"`
			CAHealthcheckEnabled                   bool        `json:"CA Healthcheck Enabled"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"IT-Root-CA"`
		SSHCA01 struct {
			ObjectType                string `json:"Object Type"`
			Version                   int    `json:"Version"`
			Name                      string `json:"Name"`
			TypeOfCA                  string `json:"Type of CA"`
			SerialNumberOctetSize     int    `json:"Serial Number Octet Size"`
			CertificateProfile        string `json:"Certificate Profile"`
			DefaultCertificateProfile string `json:"Default Certificate Profile"`
			UseAppendOnlyTable        bool   `json:"Use Append-Only Table"`
			CAToken                   struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool   `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool   `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool   `json:"Enforce Unique DN"`
			UserStorage                            bool   `json:"User Storage"`
			CertificateStorage                     bool   `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool   `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string `json:"Subject DN"`
			SignedBy                               string `json:"Signed By"`
			Validity                               string `json:"Validity"`
			UseUTF8InPolicyNoticeText              bool   `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool   `json:"LDAP DN Order"`
			CRLExpirationPeriod                    string `json:"CRL Expiration Period"`
			CRLIssueInterval                       string `json:"CRL Issue Interval"`
			CRLOverlapTime                         string `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool   `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool   `json:"Allow Changing Revocation Reason"`
			FinishUser                             bool   `json:"Finish User"`
			CAHealthcheckEnabled                   bool   `json:"CA Healthcheck Enabled"`
		} `json:"SSH-CA-01"`
		CarlSubG1 struct {
			ObjectType                    string `json:"Object Type"`
			Version                       int    `json:"Version"`
			Name                          string `json:"Name"`
			TypeOfCA                      string `json:"Type of CA"`
			SerialNumberOctetSize         int    `json:"Serial Number Octet Size"`
			PreProduceOCSPResponses       bool   `json:"Pre-produce OCSP Responses"`
			MicrosoftCACompatibleModeUsed bool   `json:"Microsoft CA Compatible Mode Used"`
			StoreResponsesOnDemand        bool   `json:"Store responses on-demand"`
			CertificateProfile            string `json:"Certificate Profile"`
			DefaultCertificateProfile     string `json:"Default Certificate Profile"`
			UseAppendOnlyTable            bool   `json:"Use Append-Only Table"`
			CAToken                       struct {
				SignatureAlgorithm    string `json:"Signature Algorithm"`
				EncryptionAlgorithm   string `json:"Encryption Algorithm"`
				CryptoToken           string `json:"Crypto Token"`
				DefaultKey            string `json:"Default Key"`
				CertificateSigningKey string `json:"Certificate Signing Key"`
				CRLSigningKey         string `json:"CRL Signing Key"`
				KeyEncryptionKey      string `json:"Key Encryption Key"`
				TestKey               string `json:"Test Key"`
				KeySequenceFormat     string `json:"Key Sequence Format"`
				KeySequence           string `json:"Key Sequence"`
			} `json:"CA Token"`
			EnforceUniquePublicKeys                bool        `json:"Enforce Unique Public Keys"`
			EnforceKeyRenewal                      bool        `json:"Enforce key renewal"`
			EnforceUniqueDN                        bool        `json:"Enforce Unique DN"`
			UserStorage                            bool        `json:"User Storage"`
			CertificateStorage                     bool        `json:"Certificate Storage"`
			AcceptRevocationsForNonExistingEntries bool        `json:"Accept Revocations for Non-Existing Entries"`
			SubjectDN                              string      `json:"Subject DN"`
			SignedBy                               string      `json:"Signed By"`
			Validity                               string      `json:"Validity"`
			SubjectAlternativeName                 string      `json:"Subject Alternative Name"`
			UseUTF8InPolicyNoticeText              bool        `json:"Use UTF-8 in Policy Notice Text"`
			LDAPDNOrder                            bool        `json:"LDAP DN Order"`
			AuthorityKeyIdUsed                     bool        `json:"Authority Key Id Used"`
			CRLNumberUsed                          bool        `json:"CRL Number Used"`
			CAIssuerURI                            []string    `json:"CA Issuer URI"`
			PartitionedCRLUsed                     bool        `json:"Partitioned CRL Used"`
			CRLExpirationPeriod                    string      `json:"CRL Expiration Period"`
			CRLIssueInterval                       string      `json:"CRL Issue Interval"`
			CRLOverlapTime                         string      `json:"CRL Overlap Time"`
			DeltaCRLPeriod                         string      `json:"Delta CRL Period"`
			GenerateCRLUponRevocation              bool        `json:"Generate CRL Upon Revocation"`
			AllowChangingRevocationReason          bool        `json:"Allow Changing Revocation Reason"`
			DefaultCRLDistributionPoint            string      `json:"Default CRL Distribution Point"`
			OCSPServiceDefaultURI                  string      `json:"OCSP Service Default URI"`
			AIACAIssuerURIs                        []string    `json:"AIA CA Issuer URIs"`
			FinishUser                             bool        `json:"Finish User"`
			CAHealthcheckEnabled                   bool        `json:"CA Healthcheck Enabled"`
			RequestProcessor                       interface{} `json:"Request Processor"`
		} `json:"Carl-Sub-G1"`
	} `json:"certification-authorities"`
	EstConfig struct {
		Device struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"device"`
		Firewall struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"firewall"`
		Estalias struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"estalias"`
		DeviceId struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"device-id"`
		EdgeCa struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"edge-ca"`
		Clientcsr struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"clientcsr"`
		Clienthttp struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"clienthttp"`
		Hayden struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"hayden"`
		Multiusera struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"multiusera"`
		Msworkstation struct {
			ObjectType                     string `json:"Object Type"`
			Version                        int    `json:"Version"`
			Name                           string `json:"Name"`
			RANameGenerationScheme         string `json:"RA Name Generation Scheme"`
			RANameGenerationDNFields       string `json:"RA Name Generation DN Fields"`
			RANameGenerationPrefix         string `json:"RA Name Generation Prefix"`
			RANameGenerationPostfix        string `json:"RA Name Generation Postfix"`
			DefaultCA                      string `json:"Default CA"`
			EndEntityProfile               string `json:"End Entity Profile"`
			CertificateProfile             string `json:"Certificate Profile"`
			RequireClientCertificate       bool   `json:"Require Client Certificate"`
			ClientUsername                 string `json:"Client Username"`
			ClientPassword                 string `json:"Client Password"`
			CertificateRenewalWithSameKeys bool   `json:"Certificate Renewal with Same Keys"`
			OperationMode                  string `json:"Operation Mode"`
			AuthenticationModule           string `json:"Authentication Module"`
			ExtractUsernameComponent       string `json:"Extract Username Component"`
			ExtractDNPartPwdComponent      string `json:"Extract DN Part Pwd Component"`
			VendorCertificateMode          bool   `json:"Vendor Certificate Mode"`
			ListOfVendorCAs                string `json:"List of Vendor CAs"`
			AllowChangeSubjectName         bool   `json:"Allow Change Subject Name"`
			SupportProxyCa                 bool   `json:"Support proxy ca"`
		} `json:"msworkstation"`
	} `json:"est-config"`
}
