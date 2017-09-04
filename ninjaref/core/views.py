from django.shortcuts import render

from .models import (
    CareerSummary,
    Ninja
)


def index(request):
    ninjas = Ninja.objects.all()
    names = ['{0} {1}'.format(n.first_name, n.last_name) for n in ninjas]
    print(names)
    return render(request, 'core/index.html', {'names': names})


def leaderboard(request):
    r = {'rating': 'speed + consistency + success'}
    leaders = CareerSummary.objects.all().extra(select=r).order_by('-rating')
    return render(request, 'core/leaderboard.html', {
        'leaders': leaders[:10],
        'men': leaders.filter(ninja__sex='M')[:10],
        'women': leaders.filter(ninja__sex='F')[:10]
    })
