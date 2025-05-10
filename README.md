# SaveUp

**SaveUp** es un backend desarrollado en **Golang** utilizando una arquitectura **hexagonal** y basada en **microservicios**. Su propósito es ayudar a los usuarios a gestionar y visualizar sus **gastos**, **ingresos** y **metas financieras** de forma organizada e inteligente.

## 🚀 Descripción del Proyecto

SaveUp permite a los usuarios:

- Registrar **gastos** e **ingresos** diarios.
- **Categorizar** y **filtrar** sus movimientos financieros.
- Crear **metas de ahorro**, especificando el monto deseado, el plazo y el propósito.
- Acceder a una sección de **análisis financiero**, que incluye:
  - Recomendaciones personalizadas.
  - Visualización mediante **gráficas interactivas**.
  - Información útil sobre cómo mejorar sus hábitos financieros.

El objetivo es empoderar a los usuarios para que tengan un mayor control sobre su economía personal.

---

## 🧱 Arquitectura

El backend de SaveUp sigue una arquitectura **Hexagonal (Ports and Adapters)**, lo que facilita la mantenibilidad, escalabilidad y pruebas del sistema. Además, cada módulo del sistema está diseñado como un **microservicio independiente**, permitiendo desplegar y escalar cada componente de forma autónoma.

### Características Técnicas

- **Lenguaje**: Go (Golang)
- **Arquitectura**: Hexagonal + Microservicios
- **Persistencia**: PostgreSQL
- **Autenticación y sesiones**
- **Métricas y monitoreo** (próximamente)
- **CI/CD** (en progreso)

---

## 📦 Estructura del Proyecto

```plaintext
saveup/
│
├── auth-service/
│   ├── cmd/ #main.go
│   ├── internal/
│   │   ├── adapter/
│   │   │   ├── http/
│   │   │   │   └── middleware/
│   │   │   └── repository/
│   │   ├── domain/
│   │   │   ├── model/
│   │   │   └── ports/
│   │   └── service/
│   ├── pkg/               # Utilidades comunes
│   └── DockerFile
├── etc-service/
├── docker-compose.yml
└── README.md

