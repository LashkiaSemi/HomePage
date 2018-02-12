class LecturesController < ApplicationController
  before_action :is_user_logged

  def index
    @lectures = Lecture.all
  end

  def new
    @lecture = Lecture.new
  end

  def create
    @lecture = Lecture.new(lecture_params)
    if @lecture.save
      flash[:success] = "登録しました"
      redirect_to lectures_url
    else
      flash[:danger] = "登録できませんでした。<br>対応していない拡張子の場合があります。"
      redirect_to root_url
    end
  end

  def destroy
    Lecture.find(params[:id]).destroy
  end

  def download
    @lecture = Lecture.find(params[:id])
    filepath = @lecture.file.current_path
    stat = File::stat(filepath)
    send_file(filepath, filename: @lecture.file.url.gsub(/.*\//,''), length: stat.size)
  end

  private
   def lecture_params
     params.require(:lecture).permit(:user_id, :title, :file, :comments, :activation)
   end
end
