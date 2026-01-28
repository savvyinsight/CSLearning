# Use SSH instead of

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

# Some Important Commands

1. For check remote URL:

```bash
git remote -v
```

2. Set remote URL:

   ```bash
   git remote set-url origin git@github.com:savvyinsight/BFT.git
   ```

3. Add remote URL:

   ```bash
   git remote add origin git@github.com:savvyinsight/BFT.git
   ```

4. Delete specific URL:

   1. Delete Local

   ```bash
   #Step 1: Switch to a different branch
   git checkout master  # or git checkout main
   
   #Step 2: Delete the local branch
   # Safe delete (recommended if branch is merged)
   git branch -d <branch-name>
   
   # Force delete (use if branch has unmerged changes you don't need)
   git branch -D <branch-name>
   ```

   2. Delete Remote

      ```bash
      # Delete a Remote Branch (on GitHub)
      git push origin --delete <branch-name>
      ```

      

# Why can't push or the push command stuck

1. Test availability first

   ```bash
   $ ssh -T -p 443 git@ssh.github.com
   ```

   **Note:**

   >Why we need to use this command to test?
   >
   >It's useful when:
   >
   >- Standard port 22 is blocked by a firewall or ISP
   >- You’re on a restricted network (corporate, school, public Wi-Fi)
   >
   >Port 443 (HTTPS) is almost always allowed, so using SSH over this port bypasses restrictions while keeping the security of SSH.
   >
   >

2. Edit : vim ~/.ssh/config

   ```bash
   Host github.com
     Hostname ssh.github.com
     Port 443
     User git
   ```

3. Then test again:

   ```bash
   ssh -T -p 443 git@ssh.github.com
   ```

   you will see : Hi ... ! You've successfully authenticated.

4. If it still don't work, **Firewall or network blocking port 443**

​	Try testing the connection to port 443:

```bash
# Test if you can reach GitHub on port 443
nc -zv ssh.github.com 443

# Or with timeout
timeout 5 nc -zv ssh.github.com 443
```



5. Another Issue : **Corporate network/proxy issues**

If you're on a corporate network:

```bash
# Check if HTTPS proxy is set
echo $https_proxy
echo $http_proxy
echo $HTTPS_PROXY
echo $HTTP_PROXY
```

or use:

```bash	
timeout 3 curl -I https://github.com
```

if curl shows HTTP/1.1 200 Connection established which means you're behind a corporate proxy.

The proxy is handling your HTTPS connections but SSH is not going through it.

so **Configure SSH to use the proxy**:

edit `~/.ssh/config`:

```bash
Host github.com
  Hostname ssh.github.com
  Port 443
  User git
  ProxyCommand nc -x 127.0.0.1:7897 %h %p
```

