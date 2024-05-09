For development

cd front
cat > .env

NUXT_MAIN_DOMAIN=http://localhost:4000
NUXT_SERVER_DOMAIN=http://localhost:4000
NUXT_BACKEND_PORT=80

For production

cat > .env

PROJECT_NAME=bestChoice

MYSQL_DATABASE=best_choice
MYSQL_ROOT_PASSWORD=root
MYSQL_USER=user
MYSQL_PASSWORD=password

NUXT_MAIN_DOMAIN=http://dual.local
NUXT_SERVER_DOMAIN=http://app:4000