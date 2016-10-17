$(document).ready(function(){

  var bh = $('body').height();
  var sh = $(window).height();
  console.log(bh, sh);
  if ((bh-60) < sh) { 
    $('.footer').addClass("pos");
  }


  $(".button").click(addArticle);
});

function addArticle(){

  var $article_title = $("<div class='content-title'><a href='#'><h3 class='font_5'>Online Marketing Hype A-Z</h3></a></div>");
  var $article_date = $("<div class='content-date'><p class='font_9 color_13'>April 5, 2016</p></div>")
  var $article_contents = $("<div class='contents'><p class='font_8 color_15'>To create your first blog post, click here and select 'Manage Posts' > New Post. Blogs are a great way to connect with your audience and keep them coming back. To really engage your site visitors we suggest you blog about subjects that are related to your site or busin...</p></div>");
  var $readmore_button = $("<div class='content-readmore'><a href='#' >Read More</a></div>");

  var $root = $("<div class='content-contain'></div>");

  $root.append($article_title);
  $root.append($article_date);
  $root.append($article_contents);
  $root.append($readmore_button);
  console.log("click");

  $("#article-contain").append($root);

}
 