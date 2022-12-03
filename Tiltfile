version_settings(constraint='>=0.30.0')
load('ext://restart_process', 'docker_build_with_restart')
load('ext://helm_remote', 'helm_remote')


# ------------ PostgreSQL ------------
postgres_password = 'LzCuZLYrNa1d2kt9rl21ZvNJPSG'

helm_remote(
  'postgresql',
  repo_url='https://charts.bitnami.com/bitnami',
  set=["auth.postgresPassword=" + postgres_password]
)

k8s_resource('postgresql', port_forwards=[5432], labels=["databases"])

# ------------ Redis ------------
redis_password = '1RNulEjFTC4q730yowi52GobJOa'

helm_remote(
  'redis',
  repo_url='https://charts.bitnami.com/bitnami',
  set=["auth.password	=" + redis_password, "architecture=standalone"]
)

k8s_resource('redis-master', port_forwards=[6379], labels=["databases"])


# ------------ Example ------------
docker_build_with_restart(
    "example-image",
    context=".",
    target="dev",
    dockerfile="services/example/Dockerfile",
    entrypoint='/app/example',
    live_update=[
        sync('services/example', '/app/services/example'),
        run('go build -mod=vendor -o /app/example ./cmd/main.go'),
    ]
)

k8s_yaml(['services/example/k8s.yaml'])

k8s_resource("example", port_forwards=8000, labels=["backend"])