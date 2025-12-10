# Guía de Instalación de Go

## 📦 Formas de Instalar Go

Hay varias formas de instalar Go. Te muestro las mejores opciones según tu sistema operativo.

## 🍎 macOS

### Opción 1: Instalador Oficial (Recomendado para principiantes)

1. **Descargar el instalador**:
   - Ve a https://go.dev/dl/
   - Descarga el archivo `.pkg` para macOS (ej: `go1.21.x.darwin-amd64.pkg`)

2. **Ejecutar el instalador**:
   ```bash
   # Abre el archivo .pkg descargado y sigue las instrucciones
   # O instala desde la línea de comandos:
   sudo installer -pkg ~/Downloads/go1.21.x.darwin-amd64.pkg -target /
   ```

3. **Verificar instalación**:
   ```bash
   go version
   # Debería mostrar: go version go1.21.x darwin/amd64
   ```

### Opción 2: Homebrew (Recomendado para desarrolladores)

```bash
# Instalar Go
brew install go

# Verificar instalación
go version
```

**Ventajas de Homebrew**:
- Fácil actualización: `brew upgrade go`
- Gestión automática de dependencias
- Integración con otros tools

### Opción 3: g (Gestor de Versiones de Go)

Si necesitas cambiar entre versiones de Go frecuentemente:

```bash
# Instalar g
go install github.com/voidint/g@latest

# Instalar una versión específica
g install 1.21.5

# Cambiar de versión
g switch 1.21.5

# Listar versiones instaladas
g ls
```

## 🐧 Linux

### Opción 1: Instalador Oficial

```bash
# Descargar
wget https://go.dev/dl/go1.21.x.linux-amd64.tar.gz

# Remover instalación anterior (si existe)
sudo rm -rf /usr/local/go

# Extraer
sudo tar -C /usr/local -xzf go1.21.x.linux-amd64.tar.gz

# Agregar a PATH (agregar a ~/.bashrc o ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin

# Recargar shell
source ~/.bashrc  # o source ~/.zshrc

# Verificar
go version
```

### Opción 2: Gestor de Paquetes

**Ubuntu/Debian**:
```bash
sudo apt update
sudo apt install golang-go
```

**Fedora/RHEL**:
```bash
sudo dnf install golang
```

**Nota**: Los repositorios pueden tener versiones más antiguas.

## 🪟 Windows

### Opción 1: Instalador MSI (Recomendado)

1. Descarga el `.msi` desde https://go.dev/dl/
2. Ejecuta el instalador
3. Go se instalará en `C:\Program Files\Go`
4. El instalador configura PATH automáticamente

### Opción 2: Chocolatey

```powershell
choco install golang
```

### Opción 3: Scoop

```powershell
scoop install go
```

## ✅ Verificar Instalación

Después de instalar, verifica que todo funcione:

```bash
# Ver versión
go version

# Ver variables de entorno importantes
go env

# Verificar GOPATH y GOROOT
go env GOPATH
go env GOROOT
```

## 🔧 Configuración Post-Instalación

### 1. Variables de Entorno Importantes

Go usa estas variables (se configuran automáticamente, pero puedes personalizarlas):

- **GOROOT**: Donde está instalado Go (normalmente `/usr/local/go` o `C:\Program Files\Go`)
- **GOPATH**: Donde Go guarda código y binarios (por defecto `~/go` o `%USERPROFILE%\go`)
- **GOBIN**: Donde se instalan binarios con `go install` (por defecto `$GOPATH/bin`)

### 2. Configurar GOPATH (Opcional)

Desde Go 1.11+, no necesitas configurar GOPATH manualmente si usas módulos (que es lo recomendado). Pero si quieres personalizarlo:

**macOS/Linux** (`~/.zshrc` o `~/.bashrc`):
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

**Windows** (PowerShell Profile):
```powershell
$env:GOPATH = "$HOME\go"
$env:PATH += ";$env:GOPATH\bin"
```

### 3. Verificar que PATH esté configurado

```bash
# Verificar que `go` está en PATH
which go  # macOS/Linux
where go  # Windows
```

## 🚀 Crear tu Primer Programa

```bash
# Crear directorio
mkdir hello-world
cd hello-world

# Inicializar módulo Go
go mod init hello-world

# Crear main.go
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
EOF

# Ejecutar
go run main.go
```

## 🔄 Actualizar Go

### macOS (Homebrew)
```bash
brew upgrade go
```

### macOS/Linux (Instalador oficial)
1. Descargar nueva versión desde https://go.dev/dl/
2. Seguir mismos pasos de instalación
3. Reemplazará la versión anterior

### Windows
1. Descargar nueva versión
2. Ejecutar instalador (reemplazará versión anterior)

## 🐛 Solución de Problemas Comunes

### "go: command not found"

**Problema**: Go no está en PATH

**Solución**:
```bash
# Verificar instalación
ls /usr/local/go/bin/go  # macOS/Linux
ls "C:\Program Files\Go\bin\go.exe"  # Windows

# Agregar a PATH manualmente si falta
export PATH=$PATH:/usr/local/go/bin  # macOS/Linux
```

### Versión incorrecta después de actualizar

**Problema**: Sigue mostrando versión antigua

**Solución**:
```bash
# Verificar qué go se está usando
which go

# Limpiar cache
go clean -cache

# Verificar PATH
echo $PATH  # macOS/Linux
$env:PATH   # Windows PowerShell
```

### Múltiples versiones instaladas

**Problema**: Conflicto entre versiones

**Solución**:
```bash
# Ver todas las instalaciones
which -a go  # macOS/Linux

# Usar versión específica o remover versiones antiguas
```

## 📚 Recomendaciones

1. **Para principiantes**: Usa el instalador oficial
2. **Para desarrolladores**: Usa Homebrew (macOS) o el instalador oficial (Linux/Windows)
3. **Para proyectos con múltiples versiones**: Usa `g` (macOS/Linux) o `gvm` (Linux)
4. **Versión mínima recomendada**: Go 1.21 o superior

## 🔗 Enlaces Útiles

- **Descargas oficiales**: https://go.dev/dl/
- **Documentación**: https://go.dev/doc/
- **Go Blog**: https://go.dev/blog/
- **Go by Example**: https://gobyexample.com/

## ✅ Checklist Post-Instalación

- [ ] `go version` funciona
- [ ] `go env` muestra configuración correcta
- [ ] Puedes crear y ejecutar un programa simple
- [ ] `go mod init` funciona
- [ ] PATH está configurado correctamente

---

**¡Listo para empezar con Go! 🚀**

