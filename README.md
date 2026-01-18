# 🐙 Hydra Security Platform

**Hydra** é uma plataforma modular de **cibersegurança ofensiva e defensiva** focada na **descoberta, análise e monitoramento de exposições de domínios, serviços e infraestruturas online**.

O projeto integra múltiplas ferramentas consolidadas de scanning, crawling e enriquecimento de dados em um **pipeline automatizado**, escalável e pronto para evoluir para um **modelo SaaS**.

> ⚠️ Este projeto é destinado a **auditorias autorizadas, pesquisa e defesa cibernética**.  
> O uso indevido é de responsabilidade exclusiva do usuário.

---

## 🎯 Objetivo do Projeto

O Hydra foi criado para oferecer **visibilidade completa da superfície de ataque**, permitindo:

- Identificar serviços expostos
- Correlacionar riscos por tipo de serviço
- Enriquecer dados técnicos com informações geográficas
- Monitorar exposições ao longo do tempo
- Automatizar análises via API

---

## 🚀 Funcionalidades

- 🔎 **Port Scanning em alta velocidade** com Masscan
- 🌐 **Coleta detalhada de banners e metadados HTTP/HTTPS** via ZGrab2
- 🧠 **Enriquecimento GeoIP** (país, ASN, organização)
- 🕷️ **Web crawling inteligente** usando Colly
- 🌪️ **Suporte a Tor e Proxies**
- 🧩 **Fuzzy Domain Detection** (descoberta de domínios similares)
- 📊 **Armazenamento estruturado e histórico**
- ⚡ **API REST para automação**
- 🛡️ **Classificação de risco por serviço**
- 📈 Endpoint de **status operacional em tempo real**

---

## 🧩 Arquitetura Geral

```text
            +-------------+
            |   Domínios  |
            +------+------+
                   |
            +------+------+
            |   Masscan   |
            +------+------+
                   |
            +------+------+
            |   ZGrab2    |
            +------+------+
                   |
        +----------+-----------+
        |                      |
+-------+------+       +-------+------+
|   GeoIP     |       |   Colly      |
+-------+------+       +-------+------+
        |                      |
        +----------+-----------+
                   |
            +------+------+
            |  Database   |
            | (MariaDB)   |
            +------+------+
                   |
            +------+------+
            |   Hydra API |
            +-------------+
```
🛠️ Tecnologias Utilizadas
Go (Golang) – Core da aplicação

Gin – Framework HTTP

Masscan – Port scanning

ZGrab2 – Coleta de serviços e banners

Colly – Web crawler

MariaDB / MySQL – Banco de dados

Tor / Proxy Chains – Anonimização

REST / JSON – Integração e automação

📂 Estrutura do Projeto
text
Copiar código
hydra/
├── cmd/
│   └── api/              # Entry point da API
├── internal/
│   ├── api/              # Controllers e rotas
│   ├── scanner/          # Integração Masscan / ZGrab2
│   ├── crawler/          # Colly
│   ├── geoip/            # Enriquecimento GeoIP
│   ├── domainfuzz/       # Fuzzy domains
│   ├── storage/          # Banco de dados
│   └── status/           # Status e métricas
├── configs/
├── scripts/
├── docs/
├── README.md
└── go.mod
⚙️ Instalação
Pré-requisitos
Go 1.21+

Masscan

ZGrab2

MariaDB ou MySQL

Tor (opcional)

Clonar o repositório
bash
Copiar código
git clone https://github.com/SEU_USUARIO/hydra.git
cd hydra
Instalar dependências
bash
Copiar código
go mod tidy
▶️ Executando a API
bash
Copiar código
go run cmd/api/main.go
A API ficará disponível em:

text
Copiar código
http://127.0.0.1:8102
🔌 Uso da API
Iniciar Scan de Domínio
bash
Copiar código
curl -X POST http://127.0.0.1:8102/v1/scan \
  -H "Content-Type: application/json" \
  -H "X-API-Key: SUA_API_KEY" \
  -d '{
    "dominio": "exemplo.com"
  }'
📊 Status Operacional
O Hydra expõe uma página de status em tempo real, exibindo:

Hosts únicos detectados

Total de exposições

Serviços mais encontrados

Classificação por risco

Estatísticas gerais da API

Ideal para monitoramento operacional e integração futura com dashboards.

🔐 Segurança
Autenticação via API Key

Controle de execução de scanners

Logs estruturados

Preparado para integração com SIEM (roadmap)

🧪 Roadmap
 Alertas (Email / Webhook / SIEM)

 Dashboard gráfico interativo

 Comparação temporal de exposições

 Modo SaaS (multi-tenant)

 Sistema de regras de risco

 Integração com Shodan / Censys

 Exportação de relatórios

🤝 Contribuição
Contribuições são bem-vindas!

Faça um fork do projeto

Crie uma branch (feature/minha-feature)

Commit suas alterações

Push para sua branch

Abra um Pull Request

📜 Licença
Este projeto é distribuído sob a licença MIT.
Consulte o arquivo LICENSE para mais detalhes.

⚠️ Aviso Legal
Este software é fornecido exclusivamente para fins educacionais, auditoria autorizada e pesquisa em segurança.
O autor não se responsabiliza por uso ilegal ou não autorizado.
