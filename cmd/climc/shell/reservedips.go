package shell

import (
	"yunion.io/x/jsonutils"
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
)

func init() {
	type NetworkReserveIPOptions struct {
		NETWORK string `help:"IP or name of network"`
		IP      string `help:"IP to reserve"`
		NOTES   string `help:"Why reserve this IP"`
	}
	R(&NetworkReserveIPOptions{}, "network-reserve-ip", "Reserve an IP address from pool", func(s *mcclient.ClientSession, args *NetworkReserveIPOptions) error {
		params := jsonutils.NewDict()
		params.Add(jsonutils.NewString(args.IP), "ip")
		params.Add(jsonutils.NewString(args.NOTES), "notes")
		net, err := modules.Networks.PerformAction(s, args.NETWORK, "reserve-ip", params)
		if err != nil {
			return err
		}
		printObject(net)
		return nil
	})

	type NetworkReleaseReservedIPOptions struct {
		NETWORK string `help:"IP or name of network"`
		IP      string `help:"IP to release"`
	}
	R(&NetworkReleaseReservedIPOptions{}, "network-release-reserved-ip", "Release a reserved IP into pool", func(s *mcclient.ClientSession, args *NetworkReleaseReservedIPOptions) error {
		params := jsonutils.NewDict()
		params.Add(jsonutils.NewString(args.IP), "ip")
		net, err := modules.Networks.PerformAction(s, args.NETWORK, "release-reserved-ip", params)
		if err != nil {
			return err
		}
		printObject(net)
		return nil
	})

	type ReservedIPListOptions struct {
		BaseListOptions
		Network string `help:"Network filter"`
	}
	R(&ReservedIPListOptions{}, "reserved-ip-list", "Show all reserved IPs for any network", func(s *mcclient.ClientSession, args *ReservedIPListOptions) error {
		params := FetchPagingParams(args.BaseListOptions)
		if len(args.Network) > 0 {
			params.Add(jsonutils.NewString(args.Network), "network")
		}
		result, err := modules.ReservedIPs.List(s, params)
		if err != nil {
			return err
		}
		printList(result, modules.ReservedIPs.GetColumns(s))
		return nil
	})

}
