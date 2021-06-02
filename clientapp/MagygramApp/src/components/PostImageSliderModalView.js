import React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from 'react-responsive-carousel';

const PostImageSliderModalView = ({media}) => {
	return (
        <React.Fragment>
            <Carousel dynamicHeight={true} showThumbs={false}>
                {media.map((media) => {
                    if(media.MediaType==="IMAGE"){
                        return (
                            <div >
                                <img width="400" height="500" className="carousel-item" src={media.Url} alt="..." />
                            </div>
                            )
                    }else if(media.MediaType==="VIDEO"){
                        return (
                            <div>
                                <video width="650" height="500" src={media.Url} alt="..." controls />
                            </div>
                        )
                    }else{
                        return <div></div>
                    }})}
            </Carousel>
        </React.Fragment>
	);
};

export default PostImageSliderModalView;
