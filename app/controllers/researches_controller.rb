class ResearchesController < ApplicationController
  before_action :is_user_logged

  def index
    @researches = Research.all
  end

  def new
    @research = Research.new
  end

  def create
    @research = Research.new(research_params)
    if @research.save
      flash[:success] = "登録しました"
      redirect_to researches_url
    else
      flash[:danger] = "登録できませんでした。<br>対応していない拡張子の場合があります。"
      redirect_to root_url
    end
  end

  def destroy
    Research.find(params[:id]).destroy
  end

  def download
    @research = Research.find(params[:id])
    filepath = @research.file.current_path
    stat = File::stat(filepath)
    send_file(filepath, filename: @research.file.url.gsub(/.*\//,''), length: stat.size)
  end

  private
    def research_params
      params.require(:research).permit(:title, :author, :file, :comments, :activation)
    end
end
