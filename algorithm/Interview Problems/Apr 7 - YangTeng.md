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
- HTTP status code
   - 2xx : Success
   - 3xx : Redirection
   - 4xx : Client Error
   - 5xx : Server Error
- 

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

2. the sences of using `Docker` technology:
Neither Kubernetes or VMware uses Docker technology by default, though they can integrate with it.
Key:
Docker -- Docker enginee.
Kubernetes -- Use CRI(Container runtime interface), container orchestrator(works with any OCI runtime)
VMware --- hypervisor , virtual machines(hardwware virtualization)

## 2.Short answer question(total 3)
1. Please describe from input `url` to response. (envolve DNS,HTTP,TCP).
2. In DevOps , what is your understanding about `CI/CD`? How do you use AI to solve CI/CD problems, at least answer `three using scene`?
3. if let you using AI to achieve "Intelligent Alarm System", how do you achieve it. Give `core component` and the `tech selection`.

`Answer:`
1. From URL to Response (7steps)
   1. Browser checks cache (DNS,HTTP cache)
   2. DNS resolution: recursive query ->IP address
   3. TCP handshake(SYN -> SYN-ACK -> ACK) + TLS(if HTTPs)
   4. HTTP request sent(GET /path HTTP/1.1)
   5. Server processes(router, business logic, DB)
   6. HTTP response(status code, body) sent back
   7. Browser renders(parse HTTP, load CSS/JS, executes)

2. CI/CD Understanding + AI Use Cases
   CI/CD :Continuous Integration(automated build/test on code commit) + Continuous Delivery/Deployment (automated release to staging/production).

   1. Flaky test prediction : ML model indentifies tests that fail non-deterministically; auto-retry or quaratine.
   2. Build failure root cause analysis: ML parses error logs, matches with historical failures, suggests fixs or assigns correct owner.
   3. Pipeline optimization : RL/ML predicts optimal parallelization & resource allocation;reduces queue time and cost.

3.   


## 3.Coding
1. Valid Parentheses . ---> **LeetCode 20**
2. Longest Palindromic Substring . --->LeetCode 5