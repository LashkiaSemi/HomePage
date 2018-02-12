class TagsController < ApplicationController
  before_action :is_user_admin

  def index
    @tag = Tag.all
  end
  
  def new
    @tag = Tag.new
  end
  
  def create
    @tag = Tag.new(tag_params)
    if @tag.save
      flash[:success] = "タグを登録しました。"
      redirect_to tag_url
    else
      flash[:danger] = "タグの登録に失敗しました。"
      render action: :new
    end
  end
  
  def destroy
    Tag.find(params[:id]).destroy
  end
end
