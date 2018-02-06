class UsersController < ApplicationController
  before_action :is_user_logged, except: [:index, :show]

  def index
    users = User.joins(:introduction).includes(:introduction)
    @seconds = users.select { |user| user.introduction.grade == 2 }
    @thirds = users.select { |user| user.introduction.grade == 3 }
    @forthes = users.select { |user| user.introduction.grade == 4 }
    @others = users.select { |user| user.introduction.grade == 0 }
  end

  def show
    @user = User.find(params[:id])
  end

  def new
    if authority_admin?
      @user = User.new
      @user.build_introduction
    else
      flash[:notice] = "権限がありません。"
      redirect_to users_url
    end
  end

  def create
    @user = User.new(user_params)
    if @user.save
      flash[:success] = "ユーザーを登録しました。"
      redirect_to users_url
    else
      flash[:danger] = "ユーザーの登録に失敗しました。"
      render action: :new
    end
  end

  def edit
    @user = User.find(params[:id])
  end

  def update
    @user = User.find(params[:id])
    if @user.update(user_params)
      flash[:success] = "ユーザーを編集しました。"
      redirect_to user_url(@user)
    else
      flash[:danger] = "ユーザーの編集に失敗しました。"
      render action: :edit
    end
  end

  def edit_pass
    @user = User.find(params[:id])
  end

  def update_pass
    @user = User.find(params[:id])
    if @user.authenticate(params[:user][:old_pass]) && @user.update(user_pass_params)
      flash[:success] = "パスワードを更新しました。"
      redirect_to user_url(@user)
    else
      flash[:danger] = "パスワードの更新に失敗しました。"
      render action: :edit_pass
    end
  end

  def destroy
    User.find(params[:id]).destroy
  end

  private
   def user_params
     params.require(:user).permit(:name, :role, :student_id, :password, :password_confirmation, introduction_attributes: [:department, :grade, :comments])
   end

   def user_pass_params
     params.require(:user).permit(:password, :password_confirmation)
   end
end
