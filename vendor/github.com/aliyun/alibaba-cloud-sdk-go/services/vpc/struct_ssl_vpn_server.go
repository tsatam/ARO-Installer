package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// SslVpnServer is a nested struct in vpc response
type SslVpnServer struct {
	InternetIp            string `json:"InternetIp" xml:"InternetIp"`
	IDaaSInstanceId       string `json:"IDaaSInstanceId" xml:"IDaaSInstanceId"`
	CreateTime            int64  `json:"CreateTime" xml:"CreateTime"`
	VpnGatewayId          string `json:"VpnGatewayId" xml:"VpnGatewayId"`
	IDaaSRegionId         string `json:"IDaaSRegionId" xml:"IDaaSRegionId"`
	Compress              bool   `json:"Compress" xml:"Compress"`
	Port                  int    `json:"Port" xml:"Port"`
	LocalSubnet           string `json:"LocalSubnet" xml:"LocalSubnet"`
	RegionId              string `json:"RegionId" xml:"RegionId"`
	Cipher                string `json:"Cipher" xml:"Cipher"`
	Connections           int    `json:"Connections" xml:"Connections"`
	SslVpnServerId        string `json:"SslVpnServerId" xml:"SslVpnServerId"`
	MaxConnections        int    `json:"MaxConnections" xml:"MaxConnections"`
	Name                  string `json:"Name" xml:"Name"`
	EnableMultiFactorAuth bool   `json:"EnableMultiFactorAuth" xml:"EnableMultiFactorAuth"`
	ClientIpPool          string `json:"ClientIpPool" xml:"ClientIpPool"`
	Proto                 string `json:"Proto" xml:"Proto"`
}