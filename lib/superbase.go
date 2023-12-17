package lib

import (
	supa "github.com/nedpals/supabase-go"
)

type SuperBaseClient struct {
	*supa.Client
}

func NewSuperBaseClient(env Env, logger Logger) SuperBaseClient {
	supabase := supa.CreateClient(env.SupabaseUrl, env.SupabaseKey)

	if supabase == nil {
		logger.Error("Không thể khởi tạo superbase")
		logger.Panic()
	}

	return SuperBaseClient{Client: supabase}
}
