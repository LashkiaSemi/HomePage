module SessionsHelper
  def log_in(user)
    session[:user_id] = user.id
  end

  def log_out
    session.delete(:user_id)
    @current_user = nil
  end

  ## ログインしているユーザーを返す
  def current_user
    @current_user ||= User.find_by(id: session[:user_id])
  end

  ## ログインしているか確認、ログインしていればtrue,していなければfalse
  def logged_in?
    !current_user.nil?
  end

  def is_user_logged
    unless logged_in?
      flash[:notice] = "ログインしてください。"
      redirect_to login_url
    end
  end
end
