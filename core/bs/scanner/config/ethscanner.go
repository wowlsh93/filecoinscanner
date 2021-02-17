/*
2021-02-10

Written by wowlsh93
*/

package config

type EthConfiguration struct {
	Simbol                         string
	Node_listen_address            string
	Start_monitoring_block         int
	Collect_account_list_path      string
	Deposit_account_list_path      string
	Deposit_account_db_path        string
	Notify_deposit_url             string
	Notify_abnormal_withdrawal_url string
	Confirmation_count             int
	Max_address_readbuffer_size    int
}
