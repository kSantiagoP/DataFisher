# DataFisher - Projeto de Enriquecimento de Dados

<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" width="200" alt="Go Logo">
  <br>
  <img src="https://i.imgur.com/6LvVyA2.png" alt="Diagrama de Arquitetura">
</p>

## 🚀 Como Rodar

```bash
docker-compose up -d --build
```

**Atenção:**  
⌛ Aguarde a inicialização completa do `db_init` antes de utilizar o sistema.  

**Problemas com imagens?**  
```bash
docker-compose up -d --force-recreate --no-deps --build nome_da_imagem
```

## 🏗️ Decisões Arquiteturais

### Estrutura do Projeto (Go Standard)
```
.
├── cmd/
│   ├── db_init/    # Inicialização do banco (migrações + dados mock)
│   ├── api/        # Publicação/consulta de jobs
│   └── worker/     # Processamento paralelo de jobs
└── internal/       # Toda a lógica da aplicação
```

**Principais escolhas:**
- **`db_init` separado**: Evita race conditions durante migrações
- **Workers isolados**: Permite escalonamento horizontal
- **Sem `pkg`**: Nenhum package precisa ser exportado

## 🧪 Testes
⚠️ **Status:** Testes manuais durante desenvolvimento  
🔜 Implementação de testes automatizados (unitários/integração) planejada

## 💡 Pontos de Melhoria

| Prioridade | Melhoria                          |
|------------|-----------------------------------|
| Alta       | Testes automatizados              |
| Alta       | Documentação Swagger              |
| Média      | Dashboard de monitoramento        |
| Média      | Circuit breaker para APIs externas|

<p align="center">
  <sub>Desenvolvido com Go • RabbitMQ • Redis • PostgreSQL</sub>
</p>