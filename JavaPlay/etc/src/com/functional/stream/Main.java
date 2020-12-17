package com.functional.stream;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) {
        ArticleStream as = new ArticleStream();
        System.out.println(Arrays.toString(as.findArticleLikeTitle("제").toArray()));
    }
}
class ArticleStream {
    Article[] articles;
    ArticleStream (){
        this.articles = new Article[]{
                new Article("제목","김작가",Arrays.asList("#김씨","#무제")),
                new Article("노인과사막","박작가",Arrays.asList("#박씨","#첫제목")),
                new Article("똑똑한제이","송작가", Arrays.asList("#송씨","#둘째 제목"))
        };
    }
    public List<Article> findArticleLikeTitle(String title){
        return Arrays.stream(this.articles)
                .filter(article -> article.getTitle().contains(title))
                .collect(Collectors.toList());
    }
}

