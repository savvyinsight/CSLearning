Firm :  扬腾创新（福建）信息科技股份有限公司
## 1.Selection
There are about 10 questions.
- 192.168.0.0/16 is given , and let you choose is private IP, public IP ...
- The command of check process.
- the scenes of using `Docker` technology.
	- VMWare
	- Kuberness
	- ...
- The reason of `High CPU`

`Answer:`
1. IP
**Categories:**
- **Public IP:** Globally unique, routable on the internet. Assigned by ISPs.
- **Private IP:** Reserved for internal networks (LAN, VMs). Not routable on the internet.
- **Loopback IP** (127.0.0.0/8): Points to your own device.
- **Link-local IP** (169.254.0.0/16): Auto-assigned when DHCP fails.
**How to differentiate:**
- **Check the first octet (or range):**
    - **Private ranges:**
        - 10.0.0.0 – 10.255.255.255 (10.0.0.0/8)
        - 172.16.0.0 – 172.31.255.255 (172.16.0.0/12)
        - 192.168.0.0 – 192.168.255.255 (192.168.0.0/16)
    - **Loopback:** 127.x.x.x
    - **Link-local:** 169.254.x.x
    - **Everything else** is typically public.
## 2.Short answer question(total 3)
1. Please describe from input `url` to response. (envolve DNS,HTTP,TCP).
2. In DevOps , what is your understanding about `CI/CD`? How do you use AI to solve CI/CD problems, at least answer `three using scene`?
3. if let you using AI to achieve "Intelligent Alarm System", how do you achieve it. Give `core component` and the `tech selection`.

## 3.Coding
1. Valid Parentheses . ---> **LeetCode 20**
2. Longest Palindromic Substring . --->LeetCode 5