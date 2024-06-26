package system

// EnvironmentData captures data for a particular environment.
type EnvironmentData struct {
	// APIHost is the hostname of an environment's public API servers.
	APIHost string

	// APIPort is the port that an environment serves the public API on.
	APIPort uint16

	// Bootstrap lists the bacalhau addresses for bootstrapping new local nodes.
	BootstrapAddresses []string

	// IPFSSwarmAddresses lists the swarm addresses of an environment's IPFS
	// nodes, for bootstrapping new local nodes.
	IPFSSwarmAddresses []string
}

// Envs is a list of environment data for various environments:
// Deprecated: stop using this, and use the config file.
var Envs = map[Environment]EnvironmentData{
	EnvironmentProd: {
		APIPort: 1234,
		APIHost: "bootstrap.production.bacalhau.org",
		BootstrapAddresses: []string{
			"/ip4/35.245.161.250/tcp/1235/p2p/QmbxGSsM6saCTyKkiWSxhJCt6Fgj7M9cns1vzYtfDbB5Ws",
			"/ip4/34.86.254.26/tcp/1235/p2p/QmeXjeQDinxm7zRiEo8ekrJdbs7585BM6j7ZeLVFrA7GPe",
			"/ip4/35.245.215.155/tcp/1235/p2p/QmPLPUUjaVE3wQNSSkxmYoaBPHVAWdjBjDYmMkWvtMZxAf",
		},
		IPFSSwarmAddresses: []string{
			"/ip4/35.245.161.250/tcp/4001/p2p/12D3KooWAQpZzf3qiNxpwizXeArGjft98ZBoMNgVNNpoWtKAvtYH",
			"/ip4/34.86.254.26/tcp/4001/p2p/12D3KooWLfFBjDo8dFe1Q4kSm8inKjPeHzmLBkQ1QAjTHocAUazK",
			"/ip4/35.245.215.155/tcp/4001/p2p/12D3KooWH3rxmhLUrpzg81KAwUuXXuqeGt4qyWRniunb5ipjemFF",
			"/ip4/34.145.201.224/tcp/4001/p2p/12D3KooWBCBZnXnNbjxqqxu2oygPdLGseEbfMbFhrkDTRjUNnZYf",
			"/ip4/35.245.41.51/tcp/4001/p2p/12D3KooWJM8j97yoDTb7B9xV1WpBXakT4Zof3aMgFuSQQH56rCXa",
		},
	},
	EnvironmentDev: {
		APIPort: 1234,
		APIHost: "bootstrap.development.bacalhau.org",
		BootstrapAddresses: []string{
			"/ip4/34.86.177.175/tcp/1235/p2p/QmfYBQ3HouX9zKcANNXbgJnpyLpTYS9nKBANw6RUQKZffu",
			"/ip4/35.245.221.171/tcp/1235/p2p/QmNjEQByyK8GiMTvnZqGyURuwXDCtzp9X6gJRKkpWfai7S",
		},
		IPFSSwarmAddresses: []string{
			"/ip4/34.86.177.175/tcp/4001/p2p/12D3KooWMSdbPzUf8WWkEcjxpCzkUfToasP9wRjFHy2iCZ6iiZdV",
			"/ip4/35.245.221.171/tcp/4001/p2p/12D3KooWRBYMhTF6MNh6eN84xcZtg6EX2wJguqEtRTNq4C7aytbu",
		},
	},
	EnvironmentStaging: {
		APIPort: 1234,
		APIHost: "bootstrap.staging.bacalhau.org",
		BootstrapAddresses: []string{
			"/ip4/34.85.228.65/tcp/1235/p2p/QmafZ9oCXCJZX9Wt1nhrGS9FVVq41qhcBRSNWCkVhz3Nvv",
			"/ip4/34.86.73.105/tcp/1235/p2p/QmVHCeiLzhFJPCyCj5S1RTAk1vBEvxd8r5A6E4HyJGQtbJ",
			"/ip4/34.150.138.100/tcp/1235/p2p/QmRr9qPTe4mU7aS9faKnWgvn1NtXt36FT8YUULRPCn2f3K",
		},
		IPFSSwarmAddresses: []string{
			"/ip4/34.85.228.65/tcp/4001/p2p/12D3KooWCWSTjjWh7SVoVv24W47z3T1Ly1tgnwZ56CCqCku5e4dS",
			"/ip4/34.86.73.105/tcp/4001/p2p/12D3KooWQuhW3LSpvhea25Zed47Z7fD5Cq2nw1xmapQ2tAUJ3q4F",
			"/ip4/34.150.138.100/tcp/4001/p2p/12D3KooWQm1T8EN8fMBz7rLviHxTGdRnohZ9nDPGbW4bfi78ckVT",
			"/ip4/35.245.247.85/tcp/4001/p2p/12D3KooWEztGEJtqtzy7th2d7cTw2iR4CQCPHFUYvj66rhh9Cf7h",
		},
	},
	EnvironmentTest: {
		APIPort: 9999,
		APIHost: "test",
		BootstrapAddresses: []string{
			"/ip4/0.0.0.0/tcp/1235/p2p/QmafZ9oCXCJZX9Wt1nhrGS9FVVq41qhcBRSNWCkVhz3Nvv",
		},
		IPFSSwarmAddresses: []string{
			"/ip4/0.0.0.0/tcp/1235/p2p/QmafZ9oCXCJZX9Wt1nhrGS9FVVq41qhcBRSNWCkVhz3Nvv",
		},
	},
}
