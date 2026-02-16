## Hugo Installation Methods on Linux

### **1. Package Manager (apt) - Easiest but outdated**
```bash
sudo apt update
sudo apt install hugo
```
✅ Simple, automatic updates via `sudo apt upgrade`
❌ Older version in repos

### **2. Manual Download (what you did)**
```bash
# Download
wget https://github.com/gohugoio/hugo/releases/download/v0.146.7/hugo_extended_0.146.7_linux-amd64.tar.gz

# Extract
tar -xzf hugo_extended_0.146.7_linux-amd64.tar.gz

# Move to PATH
sudo mv hugo /usr/local/bin/

# Verify
hugo version
```
✅ Latest version, extended features
❌ Manual updates

### **3. Snap (Universal Linux)**
```bash
sudo snap install hugo
```
✅ Auto-updates, works on all distros
❌ Slower startup

### **4. From GitHub Releases (scriptable)**
```bash
# One-liner for latest
wget -O- https://github.com/gohugoio/hugo/releases/download/v0.146.7/hugo_extended_0.146.7_linux-amd64.tar.gz | tar -xz && sudo mv hugo /usr/local/bin/
```
✅ Quick for updates

### **5. Build from source (advanced)**
```bash
git clone https://github.com/gohugoio/hugo.git
cd hugo
go install --tags extended
```
✅ Latest development version
❌ Requires Go installed

## Key Linux Concepts Used:
- **`/usr/local/bin/`** - User programs (preferred for manual installs)
- **`tar -xzf`** - Extract .tar.gz files (x=extract, z=gzip, f=file)
- **`sudo mv`** - Move with root permissions
- **`PATH`** - Directories where system looks for executables
