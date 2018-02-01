require 'test_helper'

class StaticPagesControllerTest < ActionDispatch::IntegrationTest
  test "should get index" do
    get static_pages_index_url
    assert_response :success
  end

  test "should get activitiy" do
    get static_pages_activitiy_url
    assert_response :success
  end

  test "should get equipment" do
    get static_pages_equipment_url
    assert_response :success
  end

  test "should get publication" do
    get static_pages_publication_url
    assert_response :success
  end

  test "should get job" do
    get static_pages_job_url
    assert_response :success
  end

  test "should get link" do
    get static_pages_link_url
    assert_response :success
  end

end
