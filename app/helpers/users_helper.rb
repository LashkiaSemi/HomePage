module UsersHelper
  require "securerandom"

  ## admin以上の権限か判定
  def authority_admin?
    (check_admin? || check_owner?)
  end

  ## adminか確認
  def check_admin?
    current_user.role == "admin"
  end

  ## ownerか確認
  def check_owner?
    current_user.role == "owner"
  end

  ## 学年が0だった場合卒業生と出力する関数
  def output_grade(grade)
    grade.zero? ? '卒業生' : "#{grade}年生"
  end

  ## ログインしているユーザーと指定されたユーザーが同一ユーザーか判定
  def match_login_user?(user)
    (user.id == current_user.id)
  end

  ## ランダムなパスワードを返す関数
  def secure_password
    SecureRandom.urlsafe_base64.slice(0..7)
  end

  ## admin以上の権限でなければroot_urlへリダイレクト
  def is_user_admin
    unless authority_admin?
      flash[:notice] = "権限がありません。"
      redirect_to root_url
    end
  end
end
