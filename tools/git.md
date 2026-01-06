GitHub stopped accepting passwords in 2021. Use SSH keys or generate a token in GitHub Settings → Developer settings → Personal access tokens

Switch to SSH or use a Personal Access Token:

**Quick fix:**

```bash
git remote set-url origin git@github.com:savvyinsight/CSLearning.git
```

Or use a token:
```bash
git remote set-url origin https://[YOUR_TOKEN]@github.com/savvyinsight/CSLearning.git
```
