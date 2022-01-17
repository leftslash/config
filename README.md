#### To Do
- Implement decryption
- Implement encryption utility

#### Requirements
- All config in one file
- Separate sections per Env (dev, prod)
- Formatted as space-delimitted key/val
- Comments allowed
- Sensitive data encrypted using cmd/util
- Default config file name/locations
  - /etc/default/[name]
  - ~/.[name]
  - ./.config
- Minimal command line args
  - config file
  - environment
- Command line flags override
- Utilize/wrap std flag package
