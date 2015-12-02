http://instagram-engineering.tumblr.com/post/122961624217/trending-at-instagram

* Popularity – the trend should be of interest to many people in our community.
* Novelty – the trend should be about something new. People were not posting about it before, or at least not with the same intensity.
* Timeliness – the trend should surface on Instagram while the real event is taking place.

KL divergence score - difference in probability of observing hashtag and time.
```
S(h, t) = P(h, t) * ln(P(h, t)/P’(h, t))
```
Essentially, we consider both the currently observed popularity, which is captured by P(h, t), and the novelty, computed as the ratio between our current observations and the expected baseline, P(h, t)/P’(h, t). The natural log (ln) function is used to smooth the “strength” of novelty and make it comparable to the popularity. The timeliness role is played by the t parameter, and by looking at the counters in the most recent time windows, trends will be picked up in real-time.

Turns out that while fancy things tend to have better accuracy, simple things work out well, so we ended up with selecting the maximal probability over the past week’s worth of measurements. 
* Very easy to compute and relatively low memory demand.
* Quite aggressive about suppressing non-trends with high variance.
* Quickly identifies emerging trends.

We noticed that some trends tend to disappear faster than the interest around them. For instance, the amount of posts using a hashtag that is trending at the moment will naturally decrease as soon as the event is finished. Therefore, its KL score will quickly decrease, then the hashtag won’t be trending anymore, even though people usually like to see photos and videos from the underlying event of a trend for a few hours after it is over.

In order to overcome those issues, we use an exponential decay function to define the time-to-live for previous trends, or how long we want to keep them for.
```
Sd(h, t) = SM(h) * (½)^((t - tmax)/half-life)
```
We set the decay parameter half-life to be two hours, meaning that SM(h) is halved every two hours. This way, if a hashtag or a place was a big trend a few hours ago, it may still show up as trending alongside the most recent trends.

### Grouping hashtags

We use the following notion of similarity between hashtags:

* Cooccurrences - hashtags that tend to be used together, for example #fashionweek, #dress and #model. Cooccurrences are computed by looking at recent media and counting the number of times each hashtag appears together with other hashtags.
* Edit distance - different spellings (or typos) of the same hashtag - for example #valentineday and #valentinesday - these tend not to cooccur because people rarely use both together. Spelling variations are taken care of by using Levenshtein distance.
* Topic distribution - hashtags that describe the same event - for example #gocavs, #gowarriors - these have different spelling and they are not likely to cooccur. We look at the captions used with these hashtags and run an internal tool that classifies them into a predefined set of topics. For each hashtag we look at the topic distribution (collected from all the media captions in which it appeared) and normalize it using TF-IDF.

### System

* pre-processor - the original media creation events holds metadata about the content and its creator, and in the pre-processing phase we fetch and attach to it all the data that is needed in order to apply quality filters in the next step.
* parser - extracts the hashtags or place used in a photo or video and applies quality filters. If a post violates our standards, it is not counted towards trending.
* scorer - stores time-aggregated counters for each trend. This is also where our scoring function S(h, t) is computed. Every few minutes, a line with the current value for S(h, t) is emitted.
* ranker - aggregates and ranks all the candidate trends and their trending scores.
