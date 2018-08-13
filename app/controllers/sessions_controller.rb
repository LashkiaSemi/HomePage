class SessionsController < ApplicationController
  def login_form
  end

  def login
    user = User.find_by(student_id: params[:session][:student_id])
    if user && user.authenticate(params[:session][:password])
      flash[:info] = "ログインしました。"
      log_in user
      redirect_to session[:previous_url] || root_url
    else
      flash.now[:danger] = "学籍番号かパスワードが間違っています"
      render action: :login_form
    end
  end

  def logout
    log_out
    flash[:info] = "ログアウトしました。"
    redirect_to root_url
  end
end
