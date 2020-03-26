# Gray Hat Go - Part 1

## Qualifiers

1. I'm not a professional Go developer.

2. Do not blindly run the code. Trust but verify.

3. All the content here is for demonstration purposes only.

4. Do not perform any activities against systems or applications without authorized consent.

5. It is your fault if you do not follow the aforementioned recommendations or use the code in an unethical manner.

6. Feel free to jump in and ask questions.

## Inspirations

1. Go Cypto by Example - <https://heavypackets.github.io/prezo/Go/go-crypto/> - Sabree Blackmon

2. Black Hat Go - <https://nostarch.com/blackhatgo>

## Why Go is a good fit for writing security tools - chapter introduction in BHG

1. Go is simple, user-friendly, compared to other statically typed languages.

2. Go performs more optimally compared to dynamically typed langagues.

3. Go is designed for multicore environments.

4. Go has a clean, integrated package managment system.

5. Go is is easy to cross-complile Go programs.

6. Go has a decent standard library.

7. Go programs are easily distributed since Go programs get built into a static binary.

8. Go tends to be pragmatic.

9. Go complile times are short.

## Why other languages may be better for writing security tools - chapter introduction in BHG

1. Go binaries are huge compared to binaries for C/C++.

2. Go is more verbose, less compact compared to languages like Python, Ruby.

3. Go is still a relatively new language (2009).

4. Go's package management took a while to reach prime time.

5. Something like tinygo may be better for embedded systems.

## Overview of the Penetration Testing Execution Standard

<http://www.pentest-standard.org/>

1. Pre-engagement Interactions - present and explain the tools and techniques available which aid in a successful pre-engagement step of a penetration test

2. Intelligence Gathering - provide a standard designed specifically for the pentester performing reconnaissance against a target

3. Threat Modeling - identify and prioritize potential threats and security mitigations to protect something of value, such as confidential data or intellectual property

4. Vulnerability Analysis - process of discovering flaws in systems and applications which can be leveraged by an attacker

5. Exploitation - focuses solely on establishing access to a system or resource by bypassing security restrictionsn

6. Post Exploitation - determine the value of the machine compromised and to maintain control of the machine for later uses

7. Reporting - document result of penetration test and possibly provide recommended corrections

## Developing a TCP scanning tool - chapter 5 in BHG

1. Modeled after nmap - <https://nmap.org/>

2. NMAP Primer - <https://danielmiessler.com/study/nmap/>

3. Code - <https://github.com/mikemadden42/GrayHatGo/tree/master/portscan>

```bash
$ nmap -p1-1000 scanme.nmap.org
Starting Nmap 7.80 ( https://nmap.org ) at 2020-03-24 15:02 CDT
Nmap scan report for scanme.nmap.org (45.33.32.156)
Host is up (0.065s latency).
Other addresses for scanme.nmap.org (not scanned): 2600:3c01::f03c:91ff:fe18:bb2f
Not shown: 991 closed ports
PORT    STATE    SERVICE
22/tcp  open     ssh
80/tcp  open     http
135/tcp filtered msrpc
136/tcp filtered profile
137/tcp filtered netbios-ns
138/tcp filtered netbios-dgm
139/tcp filtered netbios-ssn
445/tcp filtered microsoft-ds
593/tcp filtered http-rpc-epmap
```

```bash
$ ./portscan08 -host scanme.nmap.org
22 open
80 open
```

## Developing a DNS enumeration tool - chapter 1 in BHG

1. Modeled after amass - <https://owasp.org/www-project-amass/>

2. OWASP Amass - Users' Guide - <https://github.com/OWASP/Amass/blob/master/doc/user_guide.md>

3. Code - <https://github.com/mikemadden42/GrayHatGo/tree/master/dnsenum>

```bash
$ amass enum --passive -d microsoft.com
...
...
OWASP Amass v3.1.10                               https://github.com/OWASP/Amass
--------------------------------------------------------------------------------
34817 names discovered - api: 18320, scrape: 227, cert: 16270
```

```bash
$ ./dnsenum05 -domain microsoft.com                                                                                                                                                                                            master +
www.microsoft.com
mail.microsoft.com
remote.microsoft.com
smtp.microsoft.com
m.microsoft.com
shop.microsoft.com
ftp.microsoft.com
mail2.microsoft.com
test.microsoft.com
portal.microsoft.com
support.microsoft.com
dev.microsoft.com
bbs.microsoft.com
email.microsoft.com
...
...
```

## Future Investigation

1. Look at Go bindings for nmap - <https://github.com/Ullaakut/nmap>

2. Continue to learn Go concepts (goroutines, channels, workgroups)

3. Continue to optimize implementations
