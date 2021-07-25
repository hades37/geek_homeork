###### 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
###### 个人见解
&emsp;&emsp;在dao层遇到sql.ErrNoRows的时候不应该把该错误抛出给上层,因为对于我们的返回来说该error是一个非程序性的错误,这应当是一个正常的查询结果我们业务端是否要查不到记录作为一个错误应当是业务端处理的,而不是在数据端处理。
```golang
//这里引用Gorm的定义：
ErrRecordNotFound = errors.New("record not found")

type Errors []error

func IsRecordNotFoundError(err error) bool {
	if errs, ok := err.(Errors); ok {
		for _, err := range errs {
			if err == ErrRecordNotFound {
				return true
			}
		}
	}
	return err == ErrRecordNotFound
}

err := db.Model(&Tabel{}).First(&result).Error
	if err != nil {
        // dao层发现的recordNotFound直接忽略该错误,然后将返回值返回为nil
        // 这样在我们调用端得到nil的返回值时其实就已经知道了其实发生了recordNotfound,如何处理这样的查询结果交由上层调用端去处理.
		if gorm.IsRecordNotFoundError(err) {
			return result, false, nil
        }
    }
// 对于其它的error,比如说InvalidSQL、DipilcatedKey之类的错误我们也应当在本地处理掉(告警之类的)，而不是wrap给上层服务,因为上层服务也无法处理该错误。
```