The Frictionless data package wants to open by file or URL (must be
looking for io.Reader?   wonder if I can sub in the S3 reader?)

However, I really need something that allows people to DL the package
anyway, so I will use a simple proxy for that.

Will want some URL patterns

/id/ID   (GET on package ID)
/doc/ID  (The 303 GET redirection for the ID)

/data/ID  (the download URL for getting the data)
should get to this too on access to /id with certain accept headers

