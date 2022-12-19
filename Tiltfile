# ------------ Tilt Settings ------------
version_settings(constraint='>=0.30.0')
load('ext://restart_process', 'docker_build_with_restart')


# ------------ Service Variables ------------

# PostgresSQL
postgres_port="5432"
postgres_password="LzCuZLYrNa1d2kt9rl21ZvNJPSG"

# Redis
redis_port="6379"
redis_password="LzCuZLYrNa1d2kt9rl21ZvNJPSG"


# ReactTV/Server
server_port="8000"

# ReactTV/Frontend
frontend_port="3000"
frontend_storybook_port="6006"


# ------------ Helm Deploy ------------

reacttv_helm_dir="./deploy/chart"
reacttv_helm = helm(
  reacttv_helm_dir,
  name='reacttv-dev',
  set=[
    'server.port='+server_port,
    'frontend.port='+frontend_port,
    'frontend.storybook.port='+frontend_storybook_port,
    'postgresql.auth.password='+postgres_password,
    'redis.auth.password='+redis_password,
  ],
  )

k8s_yaml(reacttv_helm, allow_duplicates=True)


# ------------ Service Port-Forwards ------------
k8s_resource('postgresql', port_forwards=[postgres_port], labels=["databases"])
k8s_resource('redis-master', port_forwards=[redis_port], labels=["databases"])
k8s_resource("server", port_forwards=server_port, labels=["backend"])
k8s_resource("frontend", port_forwards=frontend_port, labels=["frontend"])
k8s_resource("storybook", port_forwards=frontend_storybook_port, labels=["frontend"])


# ------------ Docker Builds ------------

docker_build_with_restart(
    "server-image",
    context=".",
    target="dev",
    dockerfile="services/server/Dockerfile",
    entrypoint='/app/server server',
    live_update=[
        sync('services/server', '/app/services/server'),
        run('go build -mod=vendor -o /app/server ./cmd/main.go'),
    ]
)

docker_build(
    "frontend",
    context="./services/frontend",
    live_update=[
        sync('services/frontend', '/usr/src/app')
    ]
)

docker_build(
    "storybook",
    context="./services/frontend",
    dockerfile="./services/frontend/DockerfileStorybook",
    live_update=[
        sync('services/frontend', '/usr/src/app')
    ]
)