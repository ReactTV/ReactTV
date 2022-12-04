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


# ------------ ReactTV-Server ------------
docker_build_with_restart(
    "server-image",
    context=".",
    target="dev",
    dockerfile="services/server/Dockerfile",
    entrypoint='/app/server',
    live_update=[
        sync('services/server', '/app/services/server'),
        run('go build -mod=vendor -o /app/server ./cmd/main.go'),
    ]
)

k8s_yaml(['services/server/k8s.yaml'])

k8s_resource("server", port_forwards=8000, labels=["backend"])

# ------------ ReactTV-Frontend ------------
docker_build(
    "frontend",
    context="./services/frontend",
    live_update=[
        sync('services/frontend', '/usr/src/app')
    ]
)

k8s_resource("frontend", port_forwards="3000", labels=["frontend"])

k8s_yaml(['services/frontend/k8s.yaml'])

# ------------ ReactTV-Frontend-Storybook ------------
docker_build(
    "frontend-storybook",
    context="./services/frontend",
    dockerfile="./services/frontend/DockerfileStorybook",
    live_update=[
        sync('services/frontend', '/usr/src/app')
    ]
)

k8s_resource("frontend-storybook", port_forwards="6006", labels=["frontend"])

k8s_yaml(['services/frontend/k8sStorybook.yaml'])