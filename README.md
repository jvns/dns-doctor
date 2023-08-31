## dns doctor

**PROTOTYPE, PROBABLY HAS ONE MILLION BUGS**

You give it a domain name where something is wrong, and it runs a bunch of checks to try to figure out the reason.

**Usage**: `dns-doctor <record-type> your-domain.com` 

## How it works

It basically just runs `dig +trace` and `dig +norecurse` on your domain, parses
the output, and tries to diagnose some problems. You'll need to have `dig`
installed.

## example output

```
$ dns-doctor asdfasdfasdf.jvns.ca
Running check: no-record
  FAILURE: No record found, using nameserver 'ns-203-c.gandi.net'
Running check: cache-mismatch
  SUCCESS: Resolver doesn't have any records cached
Running check: bad-cname
  SUCCESS: No CNAME found
Running check: negative-cache
  SUCCESS: Resolver doesn't have any records cached
Running check: no-http
  SUCCESS: All A records have HTTP and HTTPS
Running check: cname-root
  SUCCESS: Skipping: this is a subdomain
```

## List of checks

### **`no-record`**

Checks to see if your authoritative nameserver actually just doesn't have a record set.

#### How it's implemented

1. Look up the record with the authoritative nameserver (the equivalent of `dig +trace some.domain.com`)
2. If there's no matching record, fail this check

#### How to fix it

DNS Doctor will give you the domain name of the nameservers where your record is set. That should give you a clue about where to fix it


### **`cache-mismatch`**

Checks to see if your local resolver has an old cached version.

#### How it's implemented

1. Look up the record with the authoritative nameserver (the equivalent of `dig +trace some.domain.com`)
2. Look up the record with the local resolver (the equivalent of `dig +norecurse some.domain.com`)
3. If the record the local resolver returns is outdated (if it's not), fail this check

#### How to fix it:

You just gotta wait! `DNS Doctor` will tell you how long you have to wait. It
figures that out by looking at the TTL (time to live) of the cached record.

### **`negative-cache`**

This is a variant of `cache-mismatch` that checks specifically for negative caching (a cached "this domain does not exist" response).

#### How it's implemented:

1. Look up the domain name with the authoritative nameserver
2. If there's no result, return success
3. Otherwise, look up the domain with local resolver, (equivalent of `dig some.domain.com`) using `+norecurse`
6. If we get an empty NOERROR result, fail the check

#### How to fix it:

You just gotta wait! `DNS Doctor` will tell you how approximately you have to wait. It
figures that out by looking at the TTL (time to live) on the domain's SOA record.

### **`bad-cname`**

Checks if your domain's CNAME doesn't resolve to anything.

#### How it's implemented:

1. Look up the domain name with the local resolver (equivalent of `dig some.domain.com`)
2. Check if it returns a CNAME record. If there's no CNAME record, return success
3. If there is a CNAME record, run `dig +trace cname.otherdomain.com`
4. If `dig +trace` reports no record for `cname.otherdomain.com`, fail the check

#### How to fix it:

You might have made a typo in your CNAME record.

### **`cname-cache-mismatch`**

1. Look up the domain name with the local resolver (equivalent of `dig some.domain.com`)
2. Check if it returns a CNAME record. If there's no CNAME record, return success
3. Run the `cache-mismatch` check on the CNAME record

### **`cname-root`**

#### How `cname-root` is implemented:

1. Check if the domain has a CNAME record
2. If it does, and if it's not a subdomain (`blah.example.com` is ok, `example.com` is not), fail the check

#### How to fix `cname-root`:

This isn't **necessarily** a problem -- your site might work just fine! (todo: explain)

### **`no-http`**

Only runs if you requested an A record check. This is a heuristic to see if you
might have made a typo in the IP address.

#### How `no-http` is implemented:

1. Look up the domain with the authoritative nameserver (the equivalent of `dig +trace some.domain.com`)
2. Try to connect to each resulting IP address on port 80 and port 443
3. If we get a "connection refused" error or if it times out, fail the check

#### How to fix this

If your server isn't a HTTP server, you can ignore this one. Otherwise there's likely an issue with the configuration of your HTTP server.

### **old-nameserver**

This checks for whether your nameservers changed recently, and if they're
cached with your local resolver. 

TODO: not implemented, not sure how to implement this.

#### How to fix `old-nameserver`:

You just gotta wait! Typically you'll have to wait up to 1-2 days after you
made the changes. DNS Doctor will tell you the TTL (in days) for the total
waiting time.
