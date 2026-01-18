# 🐙 Hydra Security Platform

**Hydra** é uma plataforma modular de **cibersegurança ofensiva e defensiva** voltada para **análise de exposição de domínios, serviços e infraestruturas online**.  
O projeto integra ferramentas consolidadas de scanning, coleta e enriquecimento de dados em um **pipeline automatizado**, escalável e pronto para uso como **SaaS**.

> ⚠️ Projeto focado em **pesquisa, auditoria e segurança**.  
> O uso indevido é de responsabilidade exclusiva do usuário.

---

## 🚀 Principais Funcionalidades

- 🔎 **Scan massivo de IPs e portas** (Masscan)
- 🌐 **Coleta detalhada de serviços HTTP/HTTPS** (ZGrab2)
- 🧠 **Enriquecimento de dados com GeoIP**
- 🕷️ **Web crawling e scraping inteligente** (Colly)
- 🌪️ **Suporte a Tor e Proxies**
- 🧩 **Fuzzy Domain Detection** (descoberta de domínios similares)
- 📊 **Armazenamento estruturado e análise histórica**
- ⚡ **API REST para automação**
- 🛡️ Classificação de risco por serviço e exposição

---

## 🧩 Arquitetura do Sistema

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
