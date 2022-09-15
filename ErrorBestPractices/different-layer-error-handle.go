package main

// https://mp.weixin.qq.com/s/EvkMQCPwg-B0fZonpwXodg
// go官网事例：https://pkg.go.dev/github.com/pkg/errors#WithMessage

/*
	一般在业务开发时，都会使用分层的方式进行各个业务间的解耦，例如：
    dao层(数据层):把访问数据库的代码封装起来，作用是封装对数据库的访问：增删改查，不涉及业务逻辑，只是达到按某个条件获得指定数据的要求。
	service层(业务逻辑层):它处理逻辑上的业务，而不去考虑具体的实现。
    controller层():controller层负责具体的业务模块流程的控制。在此层要调用service层的接口来控制业务流程。负责url映射（action）。
    view层：view层与控制层结合比较紧密，需要二者结合起来协同开发。view层主要负责前台jsp页面的显示。
 */


/*
	// controller
	if err := mode.ParamCheck(param); err != nil {
		log.Errorf("param=%+v", param)
		return errs.ErrInvalidParam
	}

	return mode.ListTestName("")

	// service
	_, err := dao.GetTestName(ctx, settleId)
		if err != nil {
		log.Errorf("GetTestName failed. err: %v", err)
		return errs.ErrDatabase
	}

	// dao
	if err != nil {
		log.Errorf("GetTestDao failed. uery: %s error(%v)", sql, err)
	}

    上面事例的问题点：
    分层开发导致的处处打印日志
	难以获取详细的堆栈关联
	根因丢失
 */

// 好的方法 Wrap erros


// 使用github.com/pkg/errors，我们需要error进行再次包装处理，这时候有三个函数可以选择（WithMessage/WithStack/Wrapf）。其次，如果需要对源错误类型进行自定义判断可以使用 Cause,可以获得最根本的错误原因。

/*
	// 新生成一个错误, 带堆栈信息
	func New(message string) error

	// 只附加新的信息
	func WithMessage(err error, message string) error

	// 只附加调用堆栈信息
	func WithStack(err error) error

	// 同时附加堆栈和信息
	func Wrapf(err error, format string, args ...interface{}) error

	// 获得最根本的错误原因
	func Cause(err error) error
 */

/*
    // Dao 层使用 Wrap 上抛错误
	if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.Wrapf(ierror.ErrNotFound, "query:%s", query)
        }
        return nil, errors.Wrapf(ierror.ErrDatabase,
            "query: %s error(%v)", query, err)
    }

    // Service 层追加信息
    bills, err := a.Dao.GetName(ctx, param)
    if err != nil {
        return result, errors.WithMessage(err, "GetName failed")
    }

    // MiddleWare 统一打印错误日志
    // 请求响应组装
	func (Format) Handle(next ihttp.MiddleFunc) ihttp.MiddleFunc {
		return func(ctx context.Context, req *http.Request, rsp *ihttp.Response) error {
			format := &format{Time: time.Now().Unix()}
			err := next(ctx, req, rsp)
			format.Data = rsp.Data
			if err != nil {
				format.Code, format.Msg = errCodes(ctx, err)
			}
			rsp.Data = format
			return nil
		}
	}

	// 获取错误码
	func errCodes(ctx context.Context, err error) (int, string) {
		if err != nil {
			log.CtxErrorf(ctx, "error: [%+v]", err)
		}
		var myError = new(erro.IError)
		if errors.As(err, &myError) {
			return myError.Code, myError.Msg
		}

		return code.ServerError, i18n.CodeMessage(code.ServerError)
	}

   // 和其他库进行协作 如果和其他库进行协作，
   // 考虑使用 errors.Wrap 或者 errors.Wrapf
   // 保存堆栈信息。同样适用 于和标准库协作的时候。
   _, err := os.Open(path)
	if err != nil {
	   return errors.Wrapf(err, "Open failed. [%s]", path)
	}
 */

/*
总结
MyError 作为全局 error 的底层实现，保存具体的错误码和错误信息；
MyError 向上返回错误时，第一次先用 Wrap 初始化堆栈，后续用 WithMessage 增加堆栈信息；
要判断 error 是否为指定的错误时，可以使用 errors.Cause 获取 root error，再进行和 sentinel error 判定；
github.com/pkg/errors 和标准库的 error 完全兼容，可以先替换、后续改造历史遗留的代码；
打印 error 的堆栈需要用%+v，而原来的%v 依旧为普通字符串方法；同时也要注意日志采集工具是否支持多行匹配；
log error 级别的打印栈，warn 和 info 可不打印堆栈；
 */
