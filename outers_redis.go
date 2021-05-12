package outers

import "github.com/go-redis/redis/v8"

type UnifiedRedisOptions struct {
	Name               string   `yaml:"name"`
	Network            string   `yaml:"network"`
	Addresses          []string `yaml:"addresses"`        // cluster only
	MaxRedirects       int      `yaml:"max_redirects"`    // cluster only
	ReadOnly           bool     `yaml:"read_only"`        // cluster only
	RouteByLatency     bool     `yaml:"route_by_latency"` // cluster only
	RouteRandomly      bool     `yaml:"route_randomly"`   // cluster only
	Address            string   `yaml:"address"`
	Database           int      `yaml:"database"`
	Username           string   `yaml:"username"`
	Password           string   `yaml:"password"`
	MaxRetries         int      `yaml:"max_retries"`
	MinRetryBackoff    Duration `yaml:"min_retry_backoff"`
	MaxRetryBackoff    Duration `yaml:"max_retry_backoff"`
	DialTimeout        Duration `yaml:"dial_timeout"`
	ReadTimeout        Duration `yaml:"read_timeout"`
	WriteTimeout       Duration `yaml:"write_timeout"`
	PoolSize           int      `yaml:"pool_size"`
	MinIdleConns       int      `yaml:"min_idle_conns"`
	MaxConnAge         Duration `yaml:"max_conn_age"`
	PoolTimeout        Duration `yaml:"pool_timeout"`
	IdleTimeout        Duration `yaml:"idle_timeout"`
	IdleCheckFrequency Duration `yaml:"idle_check_frequency"`
}

func (opts UnifiedRedisOptions) UnwrapRedisOptions() *redis.Options {
	return &redis.Options{
		Network:            opts.Network,
		Addr:               opts.Address,
		DB:                 opts.Database,
		Username:           opts.Username,
		Password:           opts.Password,
		MaxRetries:         opts.MaxRetries,
		MinRetryBackoff:    opts.MinRetryBackoff.Unwrap(),
		MaxRetryBackoff:    opts.MaxRetryBackoff.Unwrap(),
		DialTimeout:        opts.DialTimeout.Unwrap(),
		ReadTimeout:        opts.ReadTimeout.Unwrap(),
		WriteTimeout:       opts.WriteTimeout.Unwrap(),
		PoolSize:           opts.PoolSize,
		MinIdleConns:       opts.MinIdleConns,
		MaxConnAge:         opts.MaxConnAge.Unwrap(),
		PoolTimeout:        opts.PoolTimeout.Unwrap(),
		IdleTimeout:        opts.IdleTimeout.Unwrap(),
		IdleCheckFrequency: opts.IdleCheckFrequency.Unwrap(),
	}
}

func (opts UnifiedRedisOptions) UnwrapRedisClusterOptions() *redis.ClusterOptions {
	return &redis.ClusterOptions{
		Addrs:              opts.Addresses,
		MaxRedirects:       opts.MaxRedirects,
		ReadOnly:           opts.ReadOnly,
		RouteByLatency:     opts.RouteByLatency,
		RouteRandomly:      opts.RouteRandomly,
		Username:           opts.Username,
		Password:           opts.Password,
		MaxRetries:         opts.MaxRetries,
		MinRetryBackoff:    opts.MinRetryBackoff.Unwrap(),
		MaxRetryBackoff:    opts.MaxRetryBackoff.Unwrap(),
		DialTimeout:        opts.DialTimeout.Unwrap(),
		ReadTimeout:        opts.ReadTimeout.Unwrap(),
		WriteTimeout:       opts.WriteTimeout.Unwrap(),
		PoolSize:           opts.PoolSize,
		MinIdleConns:       opts.MinIdleConns,
		MaxConnAge:         opts.MaxConnAge.Unwrap(),
		PoolTimeout:        opts.PoolTimeout.Unwrap(),
		IdleTimeout:        opts.IdleTimeout.Unwrap(),
		IdleCheckFrequency: opts.IdleCheckFrequency.Unwrap(),
	}
}

func Redis(optKey ...string) (out *redis.Client, err error) {
	var opts *redis.Options
	if opts, err = RedisOptions(optKey...); err != nil {
		return
	}
	out = redis.NewClient(opts)
	return
}

func RedisOptions(optKeys ...string) (out *redis.Options, err error) {
	var opts UnifiedRedisOptions
	if err = Load(extractOptKeys(optKeys), "redis", &opts); err != nil {
		return
	}
	out = opts.UnwrapRedisOptions()
	return
}

func RedisCluster(optKey ...string) (out *redis.ClusterClient, err error) {
	var opts *redis.ClusterOptions
	if opts, err = RedisClusterOptions(optKey...); err != nil {
		return
	}
	out = redis.NewClusterClient(opts)
	return
}

func RedisClusterOptions(optKeys ...string) (out *redis.ClusterOptions, err error) {
	var opts UnifiedRedisOptions
	if err = Load(extractOptKeys(optKeys), "redis", &opts); err != nil {
		return
	}
	out = opts.UnwrapRedisClusterOptions()
	return
}
