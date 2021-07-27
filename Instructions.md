# eCatholic Coding Practical

----
## Overview
This coding practical is meant to help us get an idea of where you are at in terms of
software engineering and what you consider to be a production level of code. Your solution
may be written in whichever programming language you are most comfortable with, but should
be written idiomatically, follow established conventions, and use a common coding style.

This exercise should not take longer than 4 hours to fully complete.

----
## Instructions

Your code should perform at a minimum the following actions:

1. Download the structured data file from https://s3.amazonaws.com/ecatholic-hiring/data.csv.
2. Remove the field named 'Memo' from all records.
3. Add a field named "import_date" and populate it appropriately.
4. For any record that has an empty value, set the value of the field to the value "missing".
5. Remove any duplicate records.
6. Submit the records as JSON objects named 'PaymentRecord' to a REST API url with an API key in the 'X-API-KEY' header.
The API URL and API Key should be configurable. **NOTE:** It is expected the API URL does not point to a working
service; however, you should assume if the service was working it would return valid HTTP status codes with errors.

# Additional Information:

Your solution should include everything needed to execute it on another person's computer.
It should be at a 'production' level including tests and documentation.

Please create fork of this repository and make it private before submitting your work. Once complete, please give read
access for your private fork to jjaynes-ecatholic.
