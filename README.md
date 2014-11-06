failure
=======

Package failure helps avoid tedious if err != nil chains; and facilitates
resource cleanup.

Tell failure which named error parameter to use when recovering; and then just
try to fail.

    func(w io.Writer) (err error) {
        defer failure.Recover(&err)

        var users []User
        failure.Fail(GetUsers(&users))
        failure.Fail(Render(w, users))
        return
    }

[![Build Status](https://drone.io/github.com/dagoof/failure/status.png)](https://drone.io/github.com/dagoof/failure/latest)

https://godoc.org/github.com/dagoof/failure
