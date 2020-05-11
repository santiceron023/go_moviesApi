# movies# users

## Iniciar para desarrollo.

Para iniciar la base de datos con docker basta con ejecutar el comando: `docker-compose up [-d] mongodb`.

Nótese que el flag `-d` es opcional, y este lo que hace es correr el contenedor en daemon (background)

---

## Ejecutar build final (WIP)

Para construir el proyecto y correrlo basta con hacer `docker-compose up -d`. Esto levantará la base de datos, compilará y creará un contenedor con el ejecutable del proyecto y lo correrá en daemon (background).
