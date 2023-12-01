# 🔄 ObsidianWorker

Uma ferramenta para salvar os arquivos gerados pelo obsidian em um repositório Git.

### Requisitos

- Go instalado e configurado.
- Acesso ao repositório Git no qual deseja automatizar.

## 🚀 Como Usar

### Clonar o Repositório

```bash
git clone https://github.com/sauloalmeida/obsidian-worker.git
```

### Configuração

1. Certifique-se de ter o Go/Git instalado em seu sistema.
2. Adicione o diretório do Obsidian no arquivo main.go
```go
cmd.Dir = "" // Adicione o diretório do obsidian
```
3. Vá para pasta windows e execute o arquivo `setup.ps1` para criar a tarefa agendada no Windows.
   
```bash
cd windows 
```
```bash
.\setup.ps1
```

3. Isso criará uma tarefa agendada chamada "ObsidianWorker" que executará o programa diariamente às 09:00.