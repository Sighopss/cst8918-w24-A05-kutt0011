# Push to New GitHub Repository

## Steps to Push This Code to a New Repo

### 1. Create a New GitHub Repository
- Go to https://github.com/new
- Repository name: `cst8918-w24-A05-<your-username>` (or your preferred name)
- Make it **Public**
- **DO NOT** initialize with README, .gitignore, or license
- Click "Create repository"

### 2. Update Remote and Push

Run these commands in your terminal:

```bash
# Remove old remote
git remote remove origin

# Add your new GitHub repo as remote (replace with your actual repo URL)
git remote add origin https://github.com/YOUR-USERNAME/YOUR-REPO-NAME.git

# Push the fresh-start branch as main
git push -u origin fresh-start:main
```

### 3. Verify on GitHub
- Go to your new repository on GitHub
- Check that all files are present
- Review the commit history (should show 8 clean commits)

## Current Commit History

```
b30d17c add terraform lock file
492abce add architecture diagram placeholder
615ec87 add terraform output values
9864dcc define azure infrastructure resources
925a2db add cloud-init script for apache installation
4225a86 define terraform variables
a17ac2a add terraform providers configuration
8b8ab6b initial commit
```

## Note About Architecture Diagram

The repository currently has `a05-architecture.txt` as a placeholder.
You should create a proper `a05-architecture.png` image file using:
- draw.io (diagrams.net)
- Microsoft Visio
- Lucidchart
- Or any diagramming tool

Then add and commit it:
```bash
git add a05-architecture.png
git commit -m "add architecture diagram"
git push
```
