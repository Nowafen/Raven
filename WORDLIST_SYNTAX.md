# Wordlist Syntax for Raven

This document outlines the correct syntax for wordlist files used by the **Raven Subdomain Discovery Tool**. While Raven automatically cleans invalid entries in wordlists (e.g., removing trailing dots or invalid characters), providing a properly formatted wordlist ensures optimal performance and avoids unnecessary processing.

---

## Syntax Rules

- **File Format**: The wordlist must be a plain text file (`.txt`).
- **One Entry Per Line**: Each line must contain exactly one subdomain prefix (word).

### Allowed Characters
- Lowercase letters (`a-z`)
- Numbers (`0-9`)
- Hyphen (`-`) for multi-word subdomains (e.g., `web-mail`)

### Disallowed Characters
- No spaces, dots (`.`), slashes (`/`), or special characters (e.g., `*`, `?`, `<`, `>`)
- No uppercase letters (Raven converts to lowercase internally, but lowercase should be used to avoid confusion)

### Additional Rules
- No empty lines (ignored by Raven but should be avoided for clarity)
- No trailing dots (e.g., `admin.` is invalid; use `admin` instead)
- File must be encoded in **UTF-8** to avoid issues with non-ASCII characters

---

## Why Follow These Rules?

Raven validates and cleans wordlists automatically (e.g., removing invalid characters or trailing dots). However, using a correctly formatted wordlist:

- Reduces processing overhead  
- Prevents potential errors in edge cases  
- Ensures consistent and predictable behavior  

---

## Examples

### ✅ Valid Wordlist
```
admin
api
blog
web-mail
dashboard
login
staging
test123
```

### ❌ Invalid Wordlist (Will Be Cleaned by Raven)
```
admin.        # Trailing dot (cleaned to: admin)
web mail      # Space not allowed (skipped)
cloud..txt    # Invalid characters (skipped)
Test          # Uppercase (cleaned to: test)
api/v1        # Slash not allowed (skipped)
```

---

## Recommendations

- **Use Meaningful Subdomains**: Include common prefixes like `admin`, `api`, `blog`, `mail`, `test`.  
- **Keep It Simple**: Avoid overly complex or long entries unless necessary.  
- **Test Your Wordlist**: Run Raven with a small wordlist first to confirm it works.  
- **Default Wordlist**: If no wordlist is provided, Raven uses:  
  `/tmp/.raven/wordlist.txt`

---

## How to Create a Wordlist

1. Use a text editor (e.g., `nano`, `vim`, or **VS Code**) to create a `.txt` file.  
2. Add one subdomain prefix per line, following the rules above.  
3. Save the file with **UTF-8 encoding** and a `.txt` extension.  
4. Pass the wordlist to Raven with:  

```bash
./raven -d example.com -w mywordlist.txt
```

---

## Notes

- Raven’s wordlist validation removes invalid entries and logs the process unless `--silent` is used.  
- For large wordlists, ensure your system has enough memory to handle the file size.  
- Example wordlists are available in the Raven repository (e.g., `wordlist.txt`).  

For more information, check the Raven documentation or run:

```bash
./raven --help
```
