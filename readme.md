# 'Access is denied' when writing to hidden files

this problem occured to me when I tryed to run the tests from https://github.com/exercism/cli on my windows 10 laptop : I got an access denied error when it tries to write to the hidden configuration file.

I could isolate a test which shows the problem, I don't know if it comes from my windows 10 install, or if it is general.

# Environement :

```
C:\>ver

Microsoft Windows [version 10.0.17134.81]

C:\>go version
go version go1.10.2 windows/amd64
```

I get the same behaviour with `go version go1.8.7 windows/386` on the same computer.

# Test results :

All asserts pass but this one :

```
--- FAIL: TestWriteFile (0.02s)
        assertions.go:256:
                        Error Trace:    WritingToHiddenFile_test.go:35
                        Error:          Received unexpected error:
                                        open C:\Users\[USER NAME]\AppData\Local\Temp\write_to_hidden_file_test_204946587: Access is denied.
                        Test:           TestWriteFile
                        Messages:       Writing to the file after it's been hidden.
```

# Things I tested

 * If I pause the test with a debugger, and manually hide / unhide the file with windows explorer's file property pane, I can trigger / avoid the error
 * If I create an hidden file, reboot Windows (to make sure the OS don't keep any handler on it that mays prevent writting to the file), and try to write to the file from Go, I get the same error
 * Calling WritFile with permission 0660 or 0666 instead of 0600 doesn't change the result. 
