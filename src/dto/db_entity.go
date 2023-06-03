package model

// todo: конверт string участников групп в список

type DB_entity struct {
	Id                        int64  `csv:"id" json:"id"`
	Uid                       string `csv:"uid" json:"uid"`
	Domain                    string `csv:"domain" json:"domain"`
	Cn                        string `csv:"cn" json:"cn"`
	Department                string `csv:"department" json:"department"`
	Title                     string `csv:"title" json:"title"`
	Who                       string `csv:"who" json:"who"`
	LogonCount                string `csv:"logon_count" json:"logon_count"`
	NumLogons7                int    `csv:"num_logons7" json:"num_logons7"`
	NumShare7                 int    `csv:"num_share7" json:"num_share7"`
	NumFile7                  int    `csv:"num_file7" json:"num_file7"`
	NumAd7                    int    `csv:"num_ad7" json:"num_ad7"`
	NumN7                     int    `csv:"num_n7" json:"num_n7"`
	NumLogons14               int    `csv:"num_logons14" json:"num_logons14"`
	NumShare14                int    `csv:"num_share14" json:"num_share14"`
	NumFile14                 int    `csv:"num_file14" json:"num_file14"`
	NumAd14                   int    `csv:"num_ad14" json:"num_ad14"`
	NumN14                    int    `csv:"num_n14" json:"num_n14"`
	NumLogons30               int    `csv:"num_logons30" json:"num_logons30"`
	NumShare30                int    `csv:"num_share30" json:"num_share30"`
	NumFile30                 int    `csv:"num_file30" json:"num_file30"`
	NumAd30                   int    `csv:"num_ad30" json:"num_ad30"`
	NumN30                    int    `csv:"num_n30" json:"num_n30"`
	NumLogons150              int    `csv:"num_logons150" json:"num_logons150"`
	NumShare150               int    `csv:"num_share150" json:"num_share150"`
	NumFile150                int    `csv:"num_file150" json:"num_file150"`
	NumAd150                  int    `csv:"num_ad150" json:"num_ad150"`
	NumN150                   int    `csv:"num_n150" json:"num_n150"`
	NumLogons365              int    `csv:"num_logons365" json:"num_logons365"`
	NumShare365               int    `csv:"num_share365" json:"num_share365"`
	NumFile365                int    `csv:"num_file365" json:"num_file365"`
	NumAd365                  int    `csv:"num_ad365" json:"num_ad365"`
	NumN365                   int    `csv:"num_n365" json:"num_n365"`
	HasUserPrincipalName      bool   `csv:"has_user_principal_name" json:"has_user_principal_name"`
	HasMail                   bool   `csv:"has_mail" json:"has_mail"`
	HasPhone                  bool   `csv:"has_phone" json:"has_phone"`
	FlagDisabled              bool   `csv:"flag_disabled" json:"flag_disabled"`
	FlagLockout               bool   `csv:"flag_lockout" json:"flag_lockout"`
	FlagPasswordNotRequired   bool   `csv:"flag_password_not_required" json:"flag_password_not_required"`
	FlagPasswordCantChange    bool   `csv:"flag_password_cant_change" json:"flag_password_cant_change"`
	FlagDontExpirePassword    bool   `csv:"flag_dont_expire_password" json:"flag_dont_expire_password"`
	OwnedFiles                int    `csv:"owned_files" json:"owned_files"`
	NumMailboxes              int    `csv:"num_mailboxes" json:"num_mailboxes"`
	NumMemberOfGroups         int    `csv:"num_member_of_groups" json:"num_member_of_groups"`
	NumMemberOfIndirectGroups int    `csv:"num_member_of_indirect_groups" json:"num_member_of_indirect_groups"`
	MemberOfIndirectGroupsIds string `csv:"member_of_indirect_groups_ids" ` //json:"member_of_indirect_groups_ids"`
	MemberOfGroupsIds         string `csv:"member_of_groups_ids" `          //json:"member_of_groups_ids"`
	IsAdmin                   bool   `csv:"is_admin" json:"is_admin"`
	IsService                 bool   `csv:"is_service" json:"is_service"`
}
