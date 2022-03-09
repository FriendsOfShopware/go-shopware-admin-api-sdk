package go_shopware_admin_sdk

type Repository struct {
	ClientService

{{ range . }}
	{{ . }} *{{ . }}Repository
{{ end }}
}

func NewRepository(client ClientService) Repository {
	repo := Repository{
		ClientService: client,
	}

    {{ range . }}
	repo.{{ . }} = (*{{ . }}Repository)(&client)
	{{ end }}

	return repo
}
