# Mentor matcher product requirements doc

Mentor matcher is a website that aims to help altruistic mentors find mentees,
and vice versa.
Users sign up to the site and either post adverts, as a mentor, for mentorship
opportunities, or browse and respond to adverts, as a mentee.
Users can list a set of tags that they are experienced in and are willing to
mentor on.
Users can list a set of tags that they want to gain experience in through
being mentored.
There are two types of adverts: mentoring offer and mentorship request.
Both adverts have tags that are used by candidate mentees/mentors to find
desirable adverts.
For a mentor seeking a mentee, the advert has a body of text where they explain
who they are, what they want to mentor on. The advert also lists rough
time-slots where the mentor would be available to mentor. The advert also
includes tags that can be used to help match mentees to the advert and improve
discoverability.
For a mentee seeking a mentor, the advert has a body where the mentee describes
what they want to learn from a mentoring relationship. The advert has tags to
improve matching between mentors.
Users should be able to browse for candidate mentors through a listing page
where they can filter by keywords in advert bodies and tags associated to
listings. The same should be true for mentors browsing candidate mentee
adverts.
Users should be authenticated and have a mechanism to recover their password if
forgotten.
Users should be able to opt in to email notifications of new candidate
mentees/mentors for their adverts.

## Version 1
* User can sign up to create a new account
* User can login to existing account
* User can recover their account if they have lost their password
* User can post a mentoring listing 
* User can browse mentoring listings and filter by keywords and tags
* User can respond to an advert to initiate a dialog with some short message
  which will be emailed to the advertisement poster. The reply-to on the email
  will be the applicant's email address.
