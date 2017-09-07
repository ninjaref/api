from django.shortcuts import render

from .models import (CareerSummary, Ninja)


def index(request):
    ninjas = Ninja.objects.all()
    data = [{
        'id': n.ninja_id,
        'name': '{0} {1}'.format(n.first_name, n.last_name)
    } for n in ninjas]
    return render(request, 'core/index.html',
                  {'ninjas': data,
                   'page': 'index'})


def leaderboard(request):
    r = {'rating': 'speed + consistency + success'}
    leaders = CareerSummary.objects.all().extra(select=r).order_by('-rating')
    return render(
        request,
        'core/leaderboard.html', {
            'leaders': leaders[:10],
            'men': leaders.filter(ninja__sex='M')[:10],
            'women': leaders.filter(ninja__sex='F')[:10],
            'page': 'leaderboard'
        })


def profile(request, sender, target):
    ninja = Ninja.objects.get(ninja_id=target)
    summary = CareerSummary.objects.get(ninja_id=target)

    profile = [
        summary.speed, summary.consistency, summary.success,
        summary.speed + summary.consistency + summary.success
    ]
    profile = [round(float(n), 2) for n in profile]

    return render(request, 'core/profile.html', {
        'ninja': ninja,
        'name': '{0} {1}'.format(ninja.first_name, ninja.last_name),
        'summary': summary,
        'profileNRData': profile,
        'page': 'profile',
    })
